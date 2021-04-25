package ops

// AuthenticationType type of authentication to contact cluster orchestrator
type AuthenticationType int32

const (
	// AUTHENTICATION_TYPE_UNSPECIFIED Unspecified authentication type
	AUTHENTICATION_TYPE_UNSPECIFIED AuthenticationType = 0
	// AUTHENTICATION_CREDENTIAL authenticate using credentials
	AUTHENTICATION_CREDENTIAL AuthenticationType = 1
	// AUTHENTICATION_TYPE_APITOKEN authenticate using API token
	AUTHENTICATION_TYPE_APITOKEN AuthenticationType = 2
)

// OrchestratorType type of cluster orchestrator
type OrchestratorType int32

const (
	// ORCHESTRATOR_TYPE_UNSPECIFIED Unspecified orchestrator type
	ORCHESTRATOR_TYPE_UNSPECIFIED OrchestratorType = 0
	// ORCHESTRATOR_TYPE_RANCHER racher as cluster orchestratpr
	ORCHESTRATOR_TYPE_RANCHER OrchestratorType = 1
)

// NodeRole cluster node role
type NodeRole int32

const (
	// NODE_ROLE_UNSPECIFIED Unspecified node role
	NODE_ROLE_UNSPECIFIED NodeRole = 0
	// NODE_ROLE_SERVER to categorize server cluster nodes
	NODE_ROLE_SERVER NodeRole = 1
	// NODE_ROLE_AGENT to categorize agent cluster nodes
	NODE_ROLE_AGENT NodeRole = 2
)

// State Orchestrator object state
type State int32

const (
	// STATE_UNSPECIFIED Unspecified state
	STATE_UNSPECIFIED State = 0
	// STATE_INIT init state
	STATE_INIT State = 1
	// STATE_ONLINE online state
	STATE_ONLINE State = 2
	// STATE_ERROR error/offline state
	STATE_ERROR State = 3
)

// OrchestratorConfig holds the config to contact cluster orchestrator
type OrchestratorConfig struct {
	Type    OrchestratorType
	Rancher *RancherOrchestratorConfig // holds rancher config
}

// RancherOrchestratorConfig holds the config to contact rancher server
type RancherOrchestratorConfig struct {
	Server             string             // address of the cluster orchestration server
	Port               string             // to contact on cluster orchestration server
	AuthenticationType AuthenticationType // denotes the authenticate type to contact cluster orchestrator
	UserName           string
	Password           string
	APIToken           string // session token to access objects
}

// Metrics holds metrics status of a cluster
type Metrics struct {
	CPUPercentage    float32 // percentage of cpu used by the cluster
	PodsPercentage   float32 // percentage of pods deployed on the cluster
	MemoryPercentage float32 // percentage of memory used by the cluster
}

// ClusterConfig holds cluster configuration details
type ClusterConfig struct {
	Name        string
	ID          string
	ManifestURL string // used by the cluster to register itself with the orchestrator
}

// ClusterStatus holds cluster status details
type ClusterStatus struct {
	Name        string
	ID          string
	Nodes       []*Node // list of nodes onboarded to the cluster
	State       State   // cluster State
	Metrics     Metrics
	ErrorString string
	NodeErrors  map[string]string
}

// Node holds cluster node details
type Node struct {
	Name               string
	ID                 string
	ClusterID          string
	NodeIP             string
	ErrorString        string
	TotalCPUs          int64   // total available CPU cores
	TotalMemoryInBytes int64   // total available memory on the cluster node
	TotalPodsCapacity  int64   // total number of deployable pods
	UsedCPUPercentage  float32 // used CPU percentage on the cluster node
	UsedMemoryInBytes  int64
	UsedPods           int64
	State              State    // state of the cluster node
	Role               NodeRole // role of the cluster node
}
