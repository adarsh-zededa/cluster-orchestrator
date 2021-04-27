package rancher

import (
	"fmt"
	"strconv"

	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
	log "github.com/sirupsen/logrus"
)

func NewRancherClient(rancherConfig *ops.RancherOrchestratorConfig) ops.OrchestratorClient {
	rancherClient := &Client{
		AuthenticationType: rancherConfig.AuthenticationType,
		rancherAPIClient: apiClient{
			Server:   rancherConfig.Server,
			Port:     rancherConfig.Port,
			UserName: rancherConfig.UserName,
			Password: rancherConfig.Password,
			APIToken: rancherConfig.APIToken,
		},
	}
	return rancherClient
}

func (rc *Client) Login() error {
	loginResponse, err := rc.rancherAPIClient.loginWithCredentials()
	if err != nil {
		return fmt.Errorf("rancher Login failed. %v", err)
	}
	rc.rancherAPIClient.APIToken = loginResponse.Token
	return nil
}

func (rc *Client) VerifyTokenValidity() error {
	if err := rc.rancherAPIClient.verifyToken(); err != nil {
		return fmt.Errorf("token verification failed. %v", err)
	}
	return nil
}

func (rc *Client) CreateCluster(clusterName string) (*ops.ClusterConfig, error) {
	clusterCreateResponse, err := rc.rancherAPIClient.createCluster(clusterName)
	if err != nil {
		return nil, fmt.Errorf("cluster create fialed: %v", err)
	}

	clusterRegisterResponse, err := rc.rancherAPIClient.clusterRegister(clusterCreateResponse.ID)
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

func (rc *Client) GetClusterStatusByID(clusterID string) (*ops.ClusterStatus, error) {
	clusterStatusResponse, err := rc.rancherAPIClient.getClusterByID(clusterID)
	if err != nil {
		return nil, fmt.Errorf("get cluster by ID failed. %v", err)
	}
	clusterNodesListResponse, err := rc.rancherAPIClient.getClusterNodesByID(clusterID)
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
	if err := rc.rancherAPIClient.deleteClusterByID(clusterID); err != nil {
		return fmt.Errorf("cluster delete fialed: %v", err)
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
	log.Infof("getState: %s", rancherState)
	switch rancherState {
	case "creating", "pending", "notready", "waiting", "initializing":
		return ops.STATE_INIT
	case "running", "active":
		return ops.STATE_ONLINE
	default:
		return ops.STATE_ERROR
	}
}
