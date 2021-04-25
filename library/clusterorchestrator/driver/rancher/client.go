package rancher

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
	"github.com/rancher/norman/types"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	managementv3 "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/sirupsen/logrus"
)

func (rc *Client) Login() error {
	loginResponse, err := loginWithCredentials(rc.UserName, rc.Password, rc.Server, rc.Port)
	if err != nil {
		return fmt.Errorf("rancher Login failed. %v", err)
	}
	rc.APIToken = loginResponse.Token
	return nil
}

func (rc *Client) VerifyTokenValidity() error {
	if err := verifyToken(rc.APIToken, rc.Server, rc.Port); err != nil {
		return fmt.Errorf("token verification failed. %v", err)
	}
	return nil
}

func (rc *Client) CreateCluster(clusterName string) (*ops.ClusterConfig, error) {
	clusterCreateResponse, err := createCluster(clusterName, rc.APIToken, rc.Server, rc.Port)
	if err != nil {
		return nil, fmt.Errorf("cluster create fialed: %v", err)
	}

	clusterRegisterResponse, err := clusterRegister(clusterCreateResponse.ID, rc.APIToken, rc.Server, rc.Port)
	if err != nil {
		return nil, fmt.Errorf("cluster register fialed: %v", err)
	}

	clusterConfig := &ops.ClusterConfig{
		Name:        clusterName,
		ID:          clusterCreateResponse.ID,
		ManifestURL: clusterRegisterResponse.ManifestURL,
	}
	return clusterConfig, nil
}

func (rc *Client) GetClusterStatus(clusterID string) (*ops.ClusterStatus, error) {
	clusterStatusResponse, err := getClusterByID(clusterID, rc.APIToken, rc.Server, rc.Port)
	if err != nil {
		return nil, fmt.Errorf("get cluster by ID failed. %v", err)
	}
	clusterNodesListResponse, err := getClusterNodesByID(clusterID, rc.APIToken, rc.Server, rc.Port)
	if err != nil {
		return nil, fmt.Errorf("get cluster nodes by ID failed. %v", err)
	}

	var nodes []*ops.Node
	nodeErrors := make(map[string]string)
	if len(clusterNodesListResponse.Data) > 0 {
		for _, nodeSummary := range clusterNodesListResponse.Data {
			totalCPUs, err := strconv.Atoi(nodeSummary.Allocatable[ResourceCPU])
			if err != nil {
				return nil, fmt.Errorf("exception while parsing node total CPU. %v", err)
			}
			totalPods, err := strconv.Atoi(nodeSummary.Allocatable[ResourcePods])
			if err != nil {
				return nil, fmt.Errorf("exception while parsing node total Pods. %v", err)
			}
			totalMemoryInBytes, err := getMemoryInBytes(nodeSummary.Allocatable[ResourceMemory])
			if err != nil {
				return nil, fmt.Errorf("exception while parsing node total Memory. %v", err)
			}
			usedCPUPercentage, err := getCPUUsagePercentage(nodeSummary.Requested[ResourceCPU], totalCPUs)
			if err != nil {
				return nil, fmt.Errorf("exception while calculate node used CPU percantage. %v", err)
			}
			usedMemoryInBytes, err := getMemoryInBytes(nodeSummary.Requested[ResourceMemory])
			if err != nil {
				return nil, fmt.Errorf("exception while parsing node used Memory. %v", err)
			}
			usedPods := 0
			if len(nodeSummary.Requested[ResourcePods]) > 0 {
				usedPods, err = strconv.Atoi(nodeSummary.Requested[ResourcePods])
				if err != nil {
					return nil, fmt.Errorf("exception while parsing node used Pods. %v", err)
				}
			}
			role := ops.NODE_ROLE_UNSPECIFIED
			if nodeSummary.ControlPlane {
				role = ops.NODE_ROLE_SERVER
			} else if nodeSummary.Worker {
				role = ops.NODE_ROLE_AGENT
			}
			if nodeSummary.Transitioning == TransitionError {
				nodeErrors[nodeSummary.NodeName] = nodeSummary.TransitioningMessage
			}
			node := &ops.Node{
				Name:               nodeSummary.NodeName,
				ID:                 nodeSummary.ID,
				ClusterID:          nodeSummary.ClusterID,
				NodeIP:             nodeSummary.IPAddress,
				ErrorString:        nodeSummary.TransitioningMessage,
				TotalCPUs:          int64(totalCPUs),
				TotalMemoryInBytes: totalMemoryInBytes,
				TotalPodsCapacity:  int64(totalPods),
				UsedCPUPercentage:  usedCPUPercentage,
				UsedMemoryInBytes:  usedMemoryInBytes,
				UsedPods:           int64(usedPods),
				State:              getState(nodeSummary.State),
				Role:               role,
			}
			nodes = append(nodes, node)
		}
	}
	totalCPUs, err := strconv.Atoi(clusterStatusResponse.Allocatable[ResourceCPU])
	if err != nil {
		return nil, fmt.Errorf("exception while parsing cluster total CPU. %v", err)
	}
	usedCPUPercentage, err := getCPUUsagePercentage(clusterStatusResponse.Requested[ResourceCPU], totalCPUs)
	if err != nil {
		return nil, fmt.Errorf("exception while calculate cluster used CPU percantage. %v", err)
	}
	totalPods, err := strconv.Atoi(clusterStatusResponse.Allocatable[ResourcePods])
	if err != nil {
		return nil, fmt.Errorf("exception while parsing cluster total Pods. %v", err)
	}
	usedPodsPercentage := float32(0)
	if len(clusterStatusResponse.Requested[ResourcePods]) > 0 {
		usedPods, err := strconv.Atoi(clusterStatusResponse.Requested[ResourcePods])
		if err != nil {
			return nil, fmt.Errorf("exception while parsing cluster used Pods. %v", err)
		}
		if usedPods > 0 && totalPods > 0 {
			usedPodsPercentage = (float32(usedPods) / float32(totalPods)) * 100
		}
	}
	usedMemoryPercentage := float32(0)
	totalMemoryInBytes, err := getMemoryInBytes(clusterStatusResponse.Allocatable[ResourceMemory])
	if err != nil {
		return nil, fmt.Errorf("exception while parsing cluster total Memory. %v", err)
	}
	usedMemoryInBytes, err := getMemoryInBytes(clusterStatusResponse.Requested[ResourceMemory])
	if err != nil {
		return nil, fmt.Errorf("exception while parsing cluster used Memory. %v", err)
	}
	if totalMemoryInBytes > 0 && usedMemoryInBytes > 0 {
		usedMemoryPercentage = (float32(usedMemoryInBytes) / float32(totalMemoryInBytes)) * 100
	}
	clusterStatus := &ops.ClusterStatus{
		Name:  clusterStatusResponse.Name,
		ID:    clusterStatusResponse.ID,
		Nodes: nodes,
		State: getState(clusterStatusResponse.State),
		Metrics: ops.Metrics{
			CPUPercentage:    usedCPUPercentage,
			PodsPercentage:   usedPodsPercentage,
			MemoryPercentage: usedMemoryPercentage,
		},
		ErrorString: clusterStatusResponse.TransitioningMessage,
		NodeErrors:  nodeErrors,
	}

	return clusterStatus, nil
}

func (rc *Client) DeleteCluster(clusterID string) error {
	if err := deleteClusterByID(clusterID, rc.APIToken, rc.Server, rc.Port); err != nil {
		return fmt.Errorf("cluster delete fialed: %v", err)
	}
	return nil
}

// newAPIClient creates a new API client
func newAPIClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{Transport: tr}
}

func fireRancherAPICall(method, url, token string, body io.Reader, expectedStatusCode int) ([]byte, error) {
	client := newAPIClient()
	requent, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("exception while creating http request. METHOD: %s, URL: %s, Error: %v", method, url, err)
	}
	requent.Header.Add("Content-Type", "application/json")
	if token != "" {
		requent.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	startTime := time.Now()
	response, err := client.Do(requent)
	if err != nil {
		return nil, fmt.Errorf("rancher API call failed. METHOD: %s, URL: %s. Error: %v", method, url, err)
	}
	duration := time.Now().Sub(startTime)
	logrus.Infof("Rancher API call METHOD: %s, URL: %s, took %v to execute", method, url, duration)

	if response.StatusCode != expectedStatusCode {
		contents := make([]byte, 0)
		if response != nil && response.Body != nil {
			contents, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, fmt.Errorf("exception while parsing response for METHOD: %s, URL: %s. %v",
					method, url, err)
			}
		}
		return nil, fmt.Errorf("rancher API call failed. METHOD: %s, URL: %s, Response: %s",
			method, url, string(contents))
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("exception while readind response for METHOD: %s, URL: %s. %v",
			method, url, err)
	}
	return contents, nil
}

func loginWithCredentials(userName, password, server, port string) (*managementv3.Token, error) {
	loginRequestBody := v3.BasicLogin{
		Username: userName,
		Password: password,
		GenericLogin: v3.GenericLogin{
			TTLMillis:   DEFAULT_TOKEN_TTL_MILLISECOND,
			Description: "Zedcontrol API session",
		},
	}
	payload, err := json.Marshal(loginRequestBody)
	if err != nil {
		return nil, fmt.Errorf("exception while marshalling login request body for rancher login: %v", err)
	}
	loginURL := fmt.Sprintf(LOGIN_URL_TEMPLATE, fmt.Sprintf("%s:%s", server, port), RANCHER_API_VERSION)
	loginResponseBody, err := fireRancherAPICall(POST, loginURL, "", bytes.NewBuffer(payload), 201)
	if err != nil {
		return nil, fmt.Errorf("rancher login API call failed. URL: %s. Error: %v", loginURL, err)
	}
	var loginResponse managementv3.Token
	if err := json.Unmarshal(loginResponseBody, &loginResponse); err != nil {
		return nil, fmt.Errorf("exception while unmarshalling rancher login response %v", err)
	}
	return &loginResponse, nil
}

func verifyToken(token, server, port string) error {
	if token == "" {
		return fmt.Errorf("verify token validity failef. API Token not available")
	}
	pingURL := fmt.Sprintf(BASE_VERSIONED_URL_TEMPLATE, fmt.Sprintf("%s:%s", server, port), RANCHER_API_VERSION)
	_, err := fireRancherAPICall(GET, pingURL, token, nil, 200)
	if err != nil {
		return fmt.Errorf("rancher ping API call failed. URL: %s. Error: %v", pingURL, err)
	}
	return nil
}

func createCluster(clusterName, token, server, port string) (*managementv3.Cluster, error) {
	clusterCreateRequest := managementv3.Cluster{
		DockerRootDir: "/var/lib/docker",
		Resource: types.Resource{
			Type: "cluster",
		},
		Name: clusterName,
	}

	clusterCreatePayload, err := json.Marshal(clusterCreateRequest)
	if err != nil {
		return nil, fmt.Errorf("exception while marshalling cluster create request body: %v", err)
	}

	clusterCreateURL := fmt.Sprintf(CLUSTER_CREATE_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", server, port), RANCHER_API_VERSION)
	createResponseBody, err := fireRancherAPICall(POST, clusterCreateURL, token,
		bytes.NewBuffer(clusterCreatePayload), 201)
	if err != nil {
		return nil, fmt.Errorf("rancher cluster create API call failed. URL: %s, Error: %v",
			clusterCreateURL, err)
	}
	var clusterCreateResponse managementv3.Cluster
	if err := json.Unmarshal(createResponseBody, &clusterCreateResponse); err != nil {
		return nil, fmt.Errorf("exception while unmarshalling rancher cluster create response %v", err)
	}
	return &clusterCreateResponse, nil
}

func clusterRegister(clusterID, token, server, port string) (*managementv3.ClusterRegistrationToken, error) {
	clusterRegisterRequest := managementv3.ClusterRegistrationToken{
		Resource: types.Resource{
			Type: "clusterRegistrationToken",
		},
		ClusterID: clusterID,
	}
	clusterRegisterPayload, err := json.Marshal(clusterRegisterRequest)
	if err != nil {
		return nil, fmt.Errorf("exception while marshalling cluster register request body: %v", err)
	}
	clusterRegisterURL := fmt.Sprintf(CLUSTER_REGISTER_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", server, port), RANCHER_API_VERSION)

	registerResponseBody, err := fireRancherAPICall(POST, clusterRegisterURL, token,
		bytes.NewBuffer(clusterRegisterPayload), 201)
	if err != nil {
		return nil, fmt.Errorf("rancher cluster register API call failed. URL: %s, Response: %s",
			clusterRegisterURL, err)
	}
	var clusterRegisterResponse managementv3.ClusterRegistrationToken
	if err := json.Unmarshal(registerResponseBody, &clusterRegisterResponse); err != nil {
		return nil, fmt.Errorf("exception while unmarshalling rancher register create response %v", err)
	}
	return &clusterRegisterResponse, nil
}

func getClusterByID(clusterID, token, server, port string) (*managementv3.Cluster, error) {
	clusterStatusURL := fmt.Sprintf(CLUSTER_BY_ID_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", server, port), RANCHER_API_VERSION, clusterID)
	statusResponseBody, err := fireRancherAPICall(GET, clusterStatusURL, token, nil, 200)
	if err != nil {
		return nil, fmt.Errorf("rancher cluster get API call failed. URL: %s, Error: %v",
			clusterStatusURL, err)
	}
	var clusterStatusResponse managementv3.Cluster
	if err := json.Unmarshal(statusResponseBody, &clusterStatusResponse); err != nil {
		return nil, fmt.Errorf("exception while unmarshalling rancher cluster get response %v", err)
	}
	return &clusterStatusResponse, nil
}

func getClusterNodesByID(clusterID, token, server, port string) (*managementv3.NodeCollection, error) {
	clusterNodesURL := fmt.Sprintf(CLUSTER_NODES_LIST_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", server, port), RANCHER_API_VERSION, clusterID)
	nodesResponseBody, err := fireRancherAPICall(GET, clusterNodesURL, token, nil, 200)
	if err != nil {
		return nil, fmt.Errorf("rancher cluster nodes API call failed. URL: %s, Error: %v",
			clusterNodesURL, err)
	}
	var clusterNodesListResponse managementv3.NodeCollection
	if err := json.Unmarshal(nodesResponseBody, &clusterNodesListResponse); err != nil {
		return nil, fmt.Errorf("exception while unmarshalling rancher cluster nodes response %v", err)
	}
	return &clusterNodesListResponse, nil
}

func deleteClusterByID(clusterID, token, server, port string) error {
	clusterDeleteURL := fmt.Sprintf(CLUSTER_BY_ID_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", server, port), RANCHER_API_VERSION, clusterID)
	_, err := fireRancherAPICall(DELETE, clusterDeleteURL, token, nil, 200)
	if err != nil {
		return fmt.Errorf("rancher cluster delete API call failed. URL: %s, Error: %v",
			clusterDeleteURL, err)
	}
	return nil
}

func getCPUUsagePercentage(usedCPU string, totalCPUs int) (float32, error) {
	if len(usedCPU) == 0 || usedCPU == "0" {
		return 0, nil
	}
	cpuPercentage := float32(0)
	usedValue, err := strconv.Atoi(usedCPU[:len(usedCPU)-1])
	if err != nil {
		return 0, fmt.Errorf("exception while converting usedCPU value to int. %v", err)
	}
	usedUnit := usedCPU[len(usedCPU)-1:]
	multiplicationValue := 1
	switch usedUnit {
	case "m":
		multiplicationValue = 1000
	case "n":
		multiplicationValue = 1000000000
	default:
		return 0, fmt.Errorf("unknown CPU used unit: %s", usedUnit)
	}
	if usedValue > 0 && totalCPUs > 0 && multiplicationValue > 0 {
		cpuPercentage = (float32(usedValue) / (float32(totalCPUs) * float32(multiplicationValue))) * 100
	}
	return cpuPercentage, nil
}

func getMemoryInBytes(memory string) (int64, error) {
	if len(memory) < 2 {
		return 0, nil
	}
	value, err := strconv.Atoi(memory[:len(memory)-2])
	if err != nil {
		return 0, fmt.Errorf("exception while converting memory value to int. %v", err)
	}
	unit := memory[len(memory)-2:]
	multiplicationValue := 1
	switch unit {
	case "Ki":
		multiplicationValue = 1024
	case "Mi":
		multiplicationValue = 1048576
	case "Gi":
		multiplicationValue = 1073741824
	case "Ti":
		multiplicationValue = 1099511627776
	case "Pi":
		multiplicationValue = 1125899906842624
	default:
		return 0, fmt.Errorf("unknown memory used unit: %s", unit)
	}
	return int64(value * multiplicationValue), nil
}

func getState(rancherState string) ops.State {
	logrus.Infof("getState: %s", rancherState)
	switch rancherState {
	case "creating", "pending", "notready", "waiting", "initializing":
		return ops.STATE_INIT
	case "running", "active":
		return ops.STATE_ONLINE
	default:
		return ops.STATE_ERROR
	}
}
