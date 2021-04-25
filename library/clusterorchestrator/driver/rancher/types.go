package rancher

import (
	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
)

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

const (
	RANCHER_API_VERSION           = "v3"
	DEFAULT_TOKEN_TTL_MILLISECOND = 7200000 // 2hours
)

const (
	SERVER_URL_TEMPLATE                = "https://%s"
	BASE_VERSIONED_URL_TEMPLATE        = SERVER_URL_TEMPLATE + "/%s"
	LOGIN_URL_TEMPLATE                 = BASE_VERSIONED_URL_TEMPLATE + "-public/localProviders/local?action=login"
	CLUSTER_CREATE_URL_TEMPLATE        = BASE_VERSIONED_URL_TEMPLATE + "/clusters"
	CLUSTER_REGISTER_URL_TEMPLATE      = BASE_VERSIONED_URL_TEMPLATE + "/clusterregistrationtoken"
	CLUSTER_BY_ID_URL_TEMPLATE         = BASE_VERSIONED_URL_TEMPLATE + "/clusters/%s"
	CLUSTER_NODES_LIST_URL_TEMPLATE    = CLUSTER_BY_ID_URL_TEMPLATE + "/nodes"
	CLUSTER_NODES_METRICS_URL_TEMPLATE = SERVER_URL_TEMPLATE + "/k8s/clusters/%s/v1/metrics.k8s.io.nodes"
)


const (
	// CPU, in cores. (500m = .5 cores)
	ResourceCPU = "cpu"
	// Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	ResourceMemory = "memory"
	// Volume size, in bytes (e,g. 5Gi = 5GiB = 5 * 1024 * 1024 * 1024)
	ResourceStorage = "storage"
	// Local ephemeral storage, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	// The resource name for ResourceEphemeralStorage is alpha and it can change across releases.
	ResourceEphemeralStorage = "ephemeral-storage"
	// Number of pods
	ResourcePods = "pods"
	// TransitionError error transition state
	TransitionError = "error"
)

// Client object used to interact with rancher server
type Client struct {
	Server             string                 // address of the cluster orchestration server
	Port               string                 // to contact on cluster orchestration server
	AuthenticationType ops.AuthenticationType // denotes the authenticate type to contact cluster orchestrator
	UserName           string
	Password           string
	APIToken           string // session token to access objects
}

//// LoginRequest Rancher API login request body
//type LoginRequest struct {
//	Username    string `json:"username"`
//	Password    string `json:"password"`
//	Description string `json:"description"`
//	TTL         int64  `json:"ttl"` //API token validity duration in milliseconds
//}
//
//// LoginResponse holds rancher Login API response
//type LoginResponse struct {
//	AuthProvider    string            `json:"authProvider"`
//	BaseType        string            `json:"baseType"`
//	ClusterId       string            `json:"clusterId"`
//	Created         string            `json:"created"`
//	CreatedTS       int64             `json:"createdTS"`
//	CreatorId       string            `json:"creatorId"`
//	Current         bool              `json:"current"`
//	Description     string            `json:"description"`
//	Enabled         bool              `json:"enabled"`
//	Expired         bool              `json:"expired"`
//	ExpiresAt       string            `json:"expiresAt"`
//	GroupPrincipals string            `json:"groupPrincipals"`
//	ID              string            `json:"id"`
//	IsDerived       bool              `json:"isDerived"`
//	Labels          map[string]string `json:"labels"`
//	LastUpdateTime  string            `json:"lastUpdateTime"`
//	Links           map[string]string `json:"links"`
//	Name            string            `json:"name"`
//	Token           string            `json:"token"`
//	TTL             int64             `json:"ttl"`
//	Type            string            `json:"type"`
//	UserId          string            `json:"userId"`
//	UserPrincipal   string            `json:"userPrincipal"`
//	UUID            string            `json:"uuid"`
//}
//
//// ClusterCreateRequest Rancher cluster create request body
//type ClusterCreateRequest struct {
//	DockerRootDir           string            `json:"dockerRootDir"`
//	EnableClusterAlerting   bool              `json:"enableClusterAlerting"`
//	EnableClusterMonitoring bool              `json:"enableClusterMonitoring"`
//	EnableNetworkPolicy     bool              `json:"enableNetworkPolicy"`
//	WindowsPreferedCluster  bool              `json:"windowsPreferedCluster"`
//	Type                    string            `json:"type"`
//	Name                    string            `json:"name"`
//	Labels                  map[string]string `json:"labels"`
//}
//
////ClusterCreateResponse rancher cluster response
//type ClusterCreateResponse struct {
//	Actions                             map[string]string `json:"actions"`
//	Annotations                         map[string]string `json:"annotations"`
//	AppliedEnableNetworkPolicy          bool              `json:"appliedEnableNetworkPolicy"`
//	BaseType                            string            `json:"baseType"`
//	ClusterTemplateId                   string            `json:"clusterTemplateId"`
//	ClusterTemplateRevisionId           string            `json:"clusterTemplateRevisionId"`
//	Conditions                          []string          `json:"conditions"`
//	Created                             string            `json:"created"`
//	CreatedTS                           int64             `json:"createdTS"`
//	CreatorId                           string            `json:"creatorId"`
//	DefaultClusterRoleForProjectMembers string            `json:"defaultClusterRoleForProjectMembers"`
//	DefaultPodSecurityPolicyTemplateId  string            `json:"defaultPodSecurityPolicyTemplateId"`
//	DockerRootDir                       string            `json:"dockerRootDir"`
//	EnableClusterAlerting               bool              `json:"enableClusterAlerting"`
//	EnableClusterMonitoring             bool              `json:"enableClusterMonitoring"`
//	EnableNetworkPolicy                 bool              `json:"enableNetworkPolicy"`
//	ID                                  string            `json:"id"`
//	Internal                            bool              `json:"internal"`
//	IstioEnabled                        bool              `json:"istioEnabled"`
//	Labels                              map[string]string `json:"labels"`
//	Links                               map[string]string `json:"links"`
//	Name                                string            `json:"name"`
//	NodeCount                           int               `json:"nodeCount"`
//	NodeVersion                         int               `json:"nodeVersion"`
//	State                               string            `json:"state"`
//	Transitioning                       string            `json:"transitioning"`
//	TransitioningMessage                string            `json:"transitioningMessage"`
//	Type                                string            `json:"type"`
//	UUID                                string            `json:"uuid"`
//	WindowsPreferedCluster              bool              `json:"windowsPreferedCluster"`
//}
//
//// ClusterRegisterRequest Rancher cluster register request body
//type ClusterRegisterRequest struct {
//	Type      string `json:"type"`
//	ClusterId string `json:"clusterId"`
//}
//
//// ClusterRegisterResponse Rancher cluster register response body
//type ClusterRegisterResponse struct {
//	Annotations          map[string]string `json:"annotations"`
//	BaseType             string            `json:"baseType"`
//	ClusterId            string            `json:"clusterId"`
//	Command              string            `json:"command"`
//	Created              string            `json:"created"`
//	CreatedTS            int64             `json:"createdTS"`
//	CreatorId            string            `json:"creatorId"`
//	ID                   string            `json:"id"`
//	InsecureCommand      string            `json:"insecureCommand"`
//	Labels               map[string]string `json:"labels"`
//	Links                map[string]string `json:"links"`
//	ManifestUrl          string            `json:"manifestUrl"`
//	Name                 string            `json:"name"`
//	NamespaceId          string            `json:"namespaceId"`
//	NodeCommand          string            `json:"nodeCommand"`
//	State                string            `json:"state"`
//	Token                string            `json:"token"`
//	Transitioning        string            `json:"transitioning"`
//	TransitioningMessage string            `json:"transitioningMessage"`
//	Type                 string            `json:"type"`
//	UUID                 string            `json:"uuid"`
//	WindowsNodeCommand   string            `json:"windowsNodeCommand"`
//}
//
//// ClusterStatusResponse Rancher cluster status response body
//type ClusterStatusResponse struct {
//	Actions                             map[string]string        `json:"actions"`
//	AgentImage                          string                   `json:"agentImage"`
//	AgentImageOverride                  string                   `json:"agentImageOverride"`
//	Allocatable                         ClusterResource          `json:"allocatable"`
//	Annotations                         map[string]string        `json:"annotations"`
//	ApiEndpoint                         string                   `json:"apiEndpoint"`
//	AppliedEnableNetworkPolicy          bool                     `json:"appliedEnableNetworkPolicy"`
//	AppliedPodSecurityPolicyTemplateId  string                   `json:"appliedPodSecurityPolicyTemplateId"`
//	AppliedSpec                         ClusterAppliedSpec       `json:"appliedSpec"`
//	AuthImage                           string                   `json:"authImage"`
//	BaseType                            string                   `json:"baseType"`
//	CaCert                              string                   `json:"caCert"`
//	Capabilities                        ClusterCapability        `json:"capabilities"`
//	Capacity                            ClusterResource          `json:"capacity"`
//	ClusterTemplateId                   string                   `json:"clusterTemplateId"`
//	ClusterTemplateRevisionId           string                   `json:"clusterTemplateRevisionId"`
//	ComponentStatuses                   []ClusterComponentStatus `json:"componentStatuses"`
//	Conditions                          []ClusterConditions      `json:"conditions"`
//	Created                             string                   `json:"created"`
//	CreatedTS                           int64                    `json:"createdTS"`
//	CreatorId                           string                   `json:"creatorId"`
//	DefaultClusterRoleForProjectMembers string                   `json:"defaultClusterRoleForProjectMembers"`
//	DefaultPodSecurityPolicyTemplateId  string                   `json:"defaultPodSecurityPolicyTemplateId"`
//	Description                         string                   `json:"description"`
//	DesiredAgentImage                   string                   `json:"desiredAgentImage"`
//	DesiredAuthImage                    string                   `json:"desiredAuthImage"`
//	DockerRootDir                       string                   `json:"dockerRootDir"`
//	Driver                              string                   `json:"driver"`
//	EksStatus                           ClusterEksStatus         `json:"eksStatus"`
//	EnableClusterAlerting               bool                     `json:"enableClusterAlerting"`
//	EnableClusterMonitoring             bool                     `json:"enableClusterMonitoring"`
//	EnableNetworkPolicy                 bool                     `json:"enableNetworkPolicy"`
//	FleetWorkspaceName                  string                   `json:"fleetWorkspaceName"`
//	ID                                  string                   `json:"id"`
//	Internal                            bool                     `json:"internal"`
//	IstioEnabled                        bool                     `json:"istioEnabled"`
//	K3sConfig                           ClusterK3sConfig         `json:"k3sConfig"`
//	Labels                              map[string]string        `json:"labels"`
//	Limits                              ClusterResource          `json:"limits"`
//	Links                               map[string]string        `json:"links"`
//	LocalClusterAuthEndpoint            ClusterAuthEndpoint      `json:"localClusterAuthEndpoint"`
//	Name                                string                   `json:"name"`
//	NodeCount                           int64                    `json:"nodeCount"`
//	NodeVersion                         int                      `json:"NndeVersion"`
//	Provider                            string                   `json:"provider"`
//	Requested                           ClusterResource          `json:"requested"`
//	State                               string                   `json:"state"`
//	Transitioning                       string                   `json:"transitioning"`
//	TransitioningMessage                string                   `json:"transitioningMessage"`
//	Type                                string                   `json:"type"`
//	UUID                                string                   `json:"uuid"`
//	Version                             ClusterVersion           `json:"version"`
//	WindowsPreferedCluster              bool                     `json:"windowsPreferedCluster"`
//}
//
//// ClusterResource holds cluster resource details
//type ClusterResource struct {
//	CPU    string `json:"cpu"`
//	Memory string `json:"memory"`
//	Pods   string `json:"pods"`
//}
//
//// ClusterAppliedSpec Rancher cluster applied specifications
//type ClusterAppliedSpec struct {
//	AgentImageOverride                  string                    `json:"agentImageOverride"`
//	Answers                             ClusterAppliedSpecAnswers `json:"answers"`
//	ClusterTemplateId                   string                    `json:"clusterTemplateId"`
//	ClusterTemplateRevisionId           string                    `json:"clusterTemplateRevisionId"`
//	DefaultClusterRoleForProjectMembers string                    `json:"defaultClusterRoleForProjectMembers"`
//	DefaultPodSecurityPolicyTemplateId  string                    `json:"defaultPodSecurityPolicyTemplateId"`
//	Description                         string                    `json:"description"`
//	DesiredAgentImage                   string                    `json:"desiredAgentImage"`
//	DesiredAuthImage                    string                    `json:"desiredAuthImage"`
//	DisplayName                         string                    `json:"displayName"`
//	DockerRootDir                       string                    `json:"dockerRootDir"`
//	EnableClusterAlerting               bool                      `json:"enableClusterAlerting"`
//	EnableClusterMonitoring             bool                      `json:"enableClusterMonitoring"`
//	EnableNetworkPolicy                 string                    `json:"enableNetworkPolicy"`
//	Internal                            bool                      `json:"internal"`
//	LocalClusterAuthEndpoint            ClusterAuthEndpoint       `json:"localClusterAuthEndpoint"`
//	Type                                string                    `json:"type"`
//	WindowsPreferedCluster              bool                      `json:"windowsPreferedCluster"`
//}
//
//// ClusterAppliedSpecAnswers
//type ClusterAppliedSpecAnswers struct {
//	ClusterId string `json:"clusterId"`
//	ProjectId string `json:"projectId"`
//	Type      string `json:"type"`
//}
//
//// ClusterAuthEndpoint
//type ClusterAuthEndpoint struct {
//	Enabled bool   `json:"enabled"`
//	Type    string `json:"type"`
//}
//
//// ClusterCapability
//type ClusterCapability struct {
//	LoadBalancerCapabilities ClusterLoadBalancerCapability `json:"loadBalancerCapabilities"`
//	NodePoolScalingSupported bool                          `json:"nodePoolScalingSupported"`
//	PspEnabled               bool                          `json:"pspEnabled"`
//	Type                     string                        `json:"type"`
//}
//
//// ClusterLoadBalancerCapability
//type ClusterLoadBalancerCapability struct {
//	HealthCheckSupported bool   `json:"healthCheckSupported"`
//	Type                 string `json:"type"`
//}
//
//// ClusterComponentStatus
//type ClusterComponentStatus struct {
//	Conditions []ClusterComponentCondition `json:"conditions"`
//	Name       string                      `json:"name"`
//	Type       string                      `json:"type"`
//}
//
//// ClusterComponentCondition
//type ClusterComponentCondition struct {
//	Message string `json:"message"`
//	Status  string `json:"status"`
//	Type    string `json:"type"`
//}
//
//// ClusterConditions
//type ClusterConditions struct {
//	LastUpdateTime string `json:"lastUpdateTime"`
//	Status         string `json:"status"`
//	Type           string `json:"type"`
//}
//
//// ClusterEksStatus
//type ClusterEksStatus struct {
//	ManagedLaunchTemplateID       string `json:"managedLaunchTemplateID"`
//	ManagedLaunchTemplateVersions string `json:"managedLaunchTemplateVersions"`
//	PrivateRequiresTunnel         string `json:"privateRequiresTunnel"`
//	SecurityGroups                string `json:"securityGroups"`
//	Subnets                       string `json:"subnets"`
//	Type                          string `json:"type"`
//	UpstreamSpec                  string `json:"upstreamSpec"`
//	VirtualNetwork                string `json:"virtualNetwork"`
//}
//
//// ClusterK3sConfig
//type ClusterK3sConfig struct {
//	K3sUpgradeStrategy ClusterK3sUpgradeStrategy `json:"k3supgradeStrategy"`
//	KubernetesVersion  string                    `json:"kubernetesVersion"`
//	Type               string                    `json:"type"`
//}
//
//// ClusterK3sUpgradeStrategy
//type ClusterK3sUpgradeStrategy struct {
//	DrainServerNodes  bool   `json:"drainServerNodes"`
//	DrainWorkerNodes  bool   `json:"drainWorkerNodes"`
//	ServerConcurrency int    `json:"serverConcurrency"`
//	Type              string `json:"type"`
//	WorkerConcurrency int    `json:"workerConcurrency"`
//}
//
//// ClusterVersion response version
//type ClusterVersion struct {
//	BuildDate    string `json:"buildDate"`
//	Compiler     string `json:"compiler"`
//	GitCommit    string `json:"gitCommit"`
//	GitTreeState string `json:"gitTreeState"`
//	GitVersion   string `json:"gitVersion"`
//	GoVersion    string `json:"goVersion"`
//	Major        string `json:"major"`
//	Minor        string `json:"minor"`
//	Platform     string `json:"platform"`
//	Type         string `json:"type"`
//}
//
//// ClusterNodesListResponse holds Rancher cluster nodes list response
//type ClusterNodesListResponse struct {
//	Type         string              `json:"type"`
//	Links        map[string]string   `json:"links"`
//	CreateTypes  map[string]string   `json:"createTypes"`
//	Actions      map[string]string   `json:"actions"`
//	Pagination   Pagination          `json:"pagination"`
//	Sort         Sort                `json:"sort"`
//	Filters      NodesFilters        `json:"filters"`
//	ResourceType string              `json:"resourceType"`
//	Data         []managementv3.Node `json:"data"`
//}
//
//// Pagination for list response
//type Pagination struct {
//	Limit int64 `json:"limit"`
//	Total int64 `json:"total"`
//}
//
//// Sort
//type Sort struct {
//	Order   string            `json:"order"`
//	Reverse string            `json:"reverse"`
//	Links   map[string]string `json:"links"`
//}
//
//// NodesFilters holds nodes filters from list view
//type NodesFilters struct {
//	AppliedNodeVersion   string `json:"appliedNodeVersion"`
//	ClusterId            string `json:"clusterId"`
//	ControlPlane         string `json:"controlPlane"`
//	Created              string `json:"created"`
//	CreatorId            string `json:"creatorId"`
//	Description          string `json:"description"`
//	Etcd                 string `json:"etcd"`
//	ExternalIpAddress    string `json:"externalIpAddress"`
//	Hostname             string `json:"hostname"`
//	ID                   string `json:"id"`
//	Imported             string `json:"imported"`
//	IpAddress            string `json:"ipAddress"`
//	Name                 string `json:"name"`
//	NamespaceId          string `json:"namespaceId"`
//	NodeName             string `json:"nodeName"`
//	NodePoolId           string `json:"nodePoolId"`
//	NodeTemplateId       string `json:"nodeTemplateId"`
//	PodCidr              string `json:"podCidr"`
//	ProviderId           string `json:"providerId"`
//	Removed              string `json:"removed"`
//	RequestedHostname    string `json:"requestedHostname"`
//	ScaledownTime        string `json:"scaledownTime"`
//	SshUser              string `json:"sshUser"`
//	State                string `json:"state"`
//	Transitioning        string `json:"transitioning"`
//	TransitioningMessage string `json:"transitioningMessage"`
//	Unschedulable        string `json:"unschedulable"`
//	UUID                 string `json:"uuid"`
//	Worker               string `json:"worker"`
//}
//
//// NodeSummary
//type NodeSummary struct {
//	Actions              map[string]string `json:"actions"`
//	Allocatable          NodeResource      `json:"allocatable"`
//	Annotations          map[string]string `json:"annotations"`
//	AppliedNodeVersion   int               `json:"appliedNodeVersion"`
//	BaseType             string            `json:"baseType"`
//	Capacity             NodeResource      `json:"capacity"`
//	ClusterId            string            `json:"clusterId"`
//	Conditions           []NodeCondition   `json:"conditions"`
//	ControlPlane         bool              `json:"controlPlane"`
//	Created              string            `json:"created"`
//	CreatedTS            int64             `json:"createdTS"`
//	CreatorId            string            `json:"creatorId"`
//	CustomConfig         string            `json:"customConfig"`
//	Etcd                 bool              `json:"etcd"`
//	Hostname             string            `json:"hostname"`
//	ID                   string            `json:"id"`
//	Imported             bool              `json:"imported"`
//	Info                 NodeInfo          `json:"info"`
//	IpAddress            string            `json:"ipAddress"`
//	Labels               map[string]string `json:"labels"`
//	Links                map[string]string `json:"links"`
//	Name                 string            `json:"name"`
//	NamespaceId          string            `json:"namespaceId"`
//	NodeName             string            `json:"nodeName"`
//	NodePoolId           string            `json:"nodePoolId"`
//	NodeTemplateId       string            `json:"nodeTemplateId"`
//	PodCidr              string            `json:"podCidr"`
//	PodCidrs             []string          `json:"podCidrs"`
//	ProviderId           string            `json:"providerId"`
//	Requested            NodeResource      `json:"requested"`
//	RequestedHostname    string            `json:"requestedHostname"`
//	State                string            `json:"state"`
//	Transitioning        string            `json:"transitioning"`
//	TransitioningMessage string            `json:"transitioningMessage"`
//	Type                 string            `json:"type"`
//	Unschedulable        bool              `json:"unschedulable"`
//	UUID                 string            `json:"uuid"`
//	Worker               bool              `json:"worker"`
//}
//
//// NodeResource
//type NodeResource struct {
//	CPU              string `json:"cpu"`
//	EphemeralStorage string `json:"ephemeral-storage"`
//	Hugepages2Mi     string `json:"hugepages-2Mi"`
//	Memory           string `json:"memory"`
//	Pods             string `json:"pods"`
//}
//
//// NodeCondition
//type NodeCondition struct {
//	LastHeartbeatTime    string `json:"lastHeartbeatTime"`
//	LastHeartbeatTimeTS  int64  `json:"lastHeartbeatTimeTS"`
//	LastTransitionTime   string `json:"lastTransitionTime"`
//	LastTransitionTimeTS int64  `json:"lastTransitionTimeTS"`
//	Message              string `json:"message"`
//	Reason               string `json:"reason"`
//	Status               string `json:"status"`
//	Type                 string `json:"type"`
//}
//
//// NodeInfo
//type NodeInfo struct {
//	CPU        NodeCPU        `json:"cpu"`
//	Kubernetes NodeKubernetes `json:"kubernetes"`
//	Memory     NodeMemory     `json:"memory"`
//	OS         NodeOS         `json:"os"`
//}
//
//// NodeCPU
//type NodeCPU struct {
//	Count int `json:"count"`
//}
//
//// NodeKubernetes
//type NodeKubernetes struct {
//	KubeProxyVersion string `json:"kubeProxyVersion"`
//	KubeletVersion   string `json:"kubeletVersion"`
//}
//
//// NodeMemory
//type NodeMemory struct {
//	MemTotalKiB int64 `json:"memTotalKiB"`
//}
//
//// NodeOS
//type NodeOS struct {
//	DockerVersion   string `json:"dockerVersion"`
//	KernelVersion   string `json:"kernelVersion"`
//	OperatingSystem string `json:"operatingSystem"`
//}
//
//// ClusterNodesMetricsResponse
//type ClusterNodesMetricsResponse struct {
//	Type         string              `json:"type"`
//	Links        map[string]string   `json:"links"`
//	Actions      map[string]string   `json:"actions"`
//	ResourceType string              `json:"resourceType"`
//	Data         []NodeMetricSummary `json:"data"`
//}
//
//// NodeMetricSummary
//type NodeMetricSummary struct {
//	ID         string             `json:"id"`
//	Type       string             `json:"type"`
//	Links      map[string]string  `json:"links"`
//	ApiVersion string             `json:"apiVersion"`
//	Kind       string             `json:"kind"`
//	Metadata   NodeMetricMetadata `json:"metadata"`
//	Timestamp  string             `json:"timestamp"`
//	Usage      NodeMetricUsage    `json:"usage"`
//	Window     string             `json:"window"`
//}
//
//// NodeMetricMetadata
//type NodeMetricMetadata struct {
//	CreationTimestamp string          `json:"creationTimestamp"`
//	Name              string          `json:"name"`
//	Relationships     string          `json:"relationships"`
//	SelfLink          string          `json:"selfLink"`
//	State             NodeMetricState `json:"state"`
//}
//
//// NodeMetricState
//type NodeMetricState struct {
//	Error         bool   `json:"error"`
//	Message       string `json:"message"`
//	Name          string `json:"name"`
//	Transitioning bool   `json:"transitioning"`
//}
//
//// NodeMetricUsage
//type NodeMetricUsage struct {
//	CPU    string `json:"cpu"`
//	Memory string `json:"memory"`
//}
