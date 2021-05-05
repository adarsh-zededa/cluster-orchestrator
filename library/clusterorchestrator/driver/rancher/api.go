package rancher

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rancher/norman/types"
	apis "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	client "github.com/rancher/rancher/pkg/client/generated/management/v3"
	log "github.com/sirupsen/logrus"
)

type apiClient struct {
	Server   string // address of the cluster orchestration server
	Port     string // to contact on cluster orchestration server
	UserName string
	Password string
	APIToken string // session token to access objects
}

func (ac *apiClient) fireRancherAPI(method, url string, requestObject interface{}, respObject interface{}, expectedStatusCode int) error {
	// 1. handle request and response body
	var (
		payloadBuffer *bytes.Buffer
		err error
		request *http.Request
	)
	if requestObject != nil {
		payload, err := json.Marshal(requestObject)
		if err != nil {
			return fmt.Errorf("exception while marshalling request body: %v", err)
		}
		payloadBuffer = bytes.NewBuffer(payload)
	}
	if respObject == nil {
		respObject = &map[string]interface{}{}
	}

	// 2. create request
	apiClient := newAPIClient()
	if payloadBuffer == nil {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, payloadBuffer)
	}
	if err != nil {
		return fmt.Errorf("exception while creating http request. METHOD: %s, URL: %s, Error: %v",
			method, url, err)
	}
	request.Header.Add("Content-Type", "application/json")
	if ac.APIToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ac.APIToken))
	}

	// 3. fire the request
	startTime := time.Now()
	response, err := apiClient.Do(request)
	if err != nil {
		return fmt.Errorf("rancher API call failed. METHOD: %s, URL: %s. Error: %v", method, url, err)
	}
	duration := time.Now().Sub(startTime)
	log.Infof("Rancher API call METHOD: %s, URL: %s, took %v to execute", method, url, duration)

	// 4. process the response
	if response.StatusCode != expectedStatusCode {
		contents := make([]byte, 0)
		if response != nil && response.Body != nil {
			contents, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return fmt.Errorf("exception while parsing response for METHOD: %s, URL: %s. %v",
					method, url, err)
			}
		}
		return fmt.Errorf("rancher API call failed. METHOD: %s, URL: %s, Response: %s",
			method, url, string(contents))
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("exception while readind response for METHOD: %s, URL: %s. %v",
			method, url, err)
	}
	if err := json.Unmarshal(contents, respObject); err != nil {
		return fmt.Errorf("exception while unmarshalling response %v", err)
	}
	return nil
}
func (ac *apiClient) loginWithCredentials() (*client.Token, error) {
	loginRequestBody := apis.BasicLogin{
		Username: ac.UserName,
		Password: ac.Password,
		GenericLogin: apis.GenericLogin{
			TTLMillis:   DEFAULT_TOKEN_TTL_MILLISECOND,
			Description: "ZedControl API session",
		},
	}
	loginURL := fmt.Sprintf(LOGIN_URL_TEMPLATE, fmt.Sprintf("%s:%s", ac.Server, ac.Port))
	loginResponse := &client.Token{}

	if err := ac.fireRancherAPI(POST, loginURL, loginRequestBody, loginResponse, http.StatusCreated);
		err != nil {
		return nil, fmt.Errorf("rancher login error. METHOD %s, URL: %s. Error: %v", POST, loginURL, err)
	}
	return loginResponse, nil
}

func (ac *apiClient) verifyToken() error {
	if ac.APIToken == "" {
		return fmt.Errorf("verify token validity failed. API Token not available")
	}

	pingURL := fmt.Sprintf(BASE_VERSIONED_URL_TEMPLATE, fmt.Sprintf("%s:%s", ac.Server, ac.Port))

	if err := ac.fireRancherAPI(GET, pingURL, nil, nil, http.StatusOK); err != nil {
		return fmt.Errorf("rancher ping error. METHOD %s, URL: %s. Error: %v", GET, pingURL, err)
	}
	return nil
}

func (ac *apiClient) createCluster(clusterName string) (*client.Cluster, error) {
	clusterCreateRequest := client.Cluster{
		DockerRootDir: "/var/lib/docker",
		Resource: types.Resource{
			Type: "cluster",
		},
		Name: clusterName,
	}
	clusterCreateURL := fmt.Sprintf(CLUSTER_BASE_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", ac.Server, ac.Port))
	clusterCreateResponse := &client.Cluster{}

	if err := ac.fireRancherAPI(POST, clusterCreateURL, clusterCreateRequest,
		clusterCreateResponse, http.StatusCreated); err != nil {
		return nil, fmt.Errorf("rancher cluster error. METHOD %s, URL: %s, Error: %v",
			POST, clusterCreateURL, err)
	}
	return clusterCreateResponse, nil
}

func (ac *apiClient) clusterRegister(clusterID string) (*client.ClusterRegistrationToken, error) {
	clusterRegisterRequest := client.ClusterRegistrationToken{
		Resource: types.Resource{
			Type: "clusterRegistrationToken",
		},
		ClusterID: clusterID,
	}
	clusterRegisterURL := fmt.Sprintf(CLUSTER_REGISTER_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", ac.Server, ac.Port))
	clusterRegisterResponse := &client.ClusterRegistrationToken{}

	if err := ac.fireRancherAPI(POST, clusterRegisterURL, clusterRegisterRequest,
		clusterRegisterResponse, http.StatusCreated); err != nil {
		return nil, fmt.Errorf("rancher cluster register error. METHOD %s, URL: %s, Response: %s",
			POST, clusterRegisterURL, err)
	}
	return clusterRegisterResponse, nil
}

func (ac *apiClient) getClusterByID(clusterID string) (*client.Cluster, error) {
	clusterStatusURL := fmt.Sprintf(CLUSTER_BY_ID_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", ac.Server, ac.Port), clusterID)
	clusterStatusResponse := &client.Cluster{}

	if err := ac.fireRancherAPI(GET, clusterStatusURL, nil, clusterStatusResponse,
		http.StatusOK); err != nil {
		return nil, fmt.Errorf("rancher cluster get by id error. METHOD %s,  URL: %s, Error: %v",
			GET, clusterStatusURL, err)
	}
	return clusterStatusResponse, nil
}

func (ac *apiClient) getClusterList() (*client.ClusterCollection, error) {
	clusterListURL := fmt.Sprintf(CLUSTER_BASE_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", ac.Server, ac.Port))
	clusterListResponse := &client.ClusterCollection{}

	if err := ac.fireRancherAPI(GET, clusterListURL, nil, clusterListResponse,
		http.StatusOK); err != nil {
		return nil, fmt.Errorf("rancher list cluster error. METHOD %s, URL: %s, Error: %v",
			GET, clusterListURL, err)
	}
	return clusterListResponse, nil
}

func (ac *apiClient) getNodesList() (*client.NodeCollection, error) {
	nodeListURL := fmt.Sprintf(NODE_BASE_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", ac.Server, ac.Port))
	nodeListResponse := &client.NodeCollection{}

	if err := ac.fireRancherAPI(GET, nodeListURL, nil, nodeListResponse,
		http.StatusOK); err != nil {
		return nil, fmt.Errorf("rancher list nodes error. METHOD %s, URL: %s, Error: %v",
			GET, nodeListURL, err)
	}
	return nodeListResponse, nil
}

func (ac *apiClient) getClusterNodesByID(clusterID string) (*client.NodeCollection, error) {
	clusterNodesURL := fmt.Sprintf(CLUSTER_NODES_LIST_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", ac.Server, ac.Port), clusterID)
	clusterNodesListResponse := &client.NodeCollection{}

	if err := ac.fireRancherAPI(GET, clusterNodesURL, nil, clusterNodesListResponse,
		http.StatusOK); err != nil {
		return nil, fmt.Errorf("rancher list cluster nodes error. METHOD %s, URL: %s, Error: %v",
			GET, clusterNodesURL, err)
	}
	return clusterNodesListResponse, nil
}

func (ac *apiClient) deleteClusterByID(clusterID string) error {
	clusterDeleteURL := fmt.Sprintf(CLUSTER_BY_ID_URL_TEMPLATE,
		fmt.Sprintf("%s:%s", ac.Server, ac.Port), clusterID)
	if err := ac.fireRancherAPI(DELETE, clusterDeleteURL, nil, nil,
		http.StatusOK); err != nil {
		return fmt.Errorf("rancher cluster delete error. METHOD %s, URL: %s, Error: %v",
			DELETE, clusterDeleteURL, err)
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
