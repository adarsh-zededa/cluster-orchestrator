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
	SERVER_URL_TEMPLATE             = "https://%s"
	BASE_VERSIONED_URL_TEMPLATE     = SERVER_URL_TEMPLATE + "/" + RANCHER_API_VERSION
	LOGIN_URL_TEMPLATE              = BASE_VERSIONED_URL_TEMPLATE + "-public/localProviders/local?action=login"
	CLUSTER_BASE_URL_TEMPLATE       = BASE_VERSIONED_URL_TEMPLATE + "/clusters"
	NODE_BASE_URL_TEMPLATE          = BASE_VERSIONED_URL_TEMPLATE + "/nodes"
	CLUSTER_REGISTER_URL_TEMPLATE   = BASE_VERSIONED_URL_TEMPLATE + "/clusterregistrationtoken"
	CLUSTER_BY_ID_URL_TEMPLATE      = CLUSTER_BASE_URL_TEMPLATE + "/%s"
	CLUSTER_NODES_LIST_URL_TEMPLATE = CLUSTER_BY_ID_URL_TEMPLATE + "/nodes"
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
	// error transition state
	TransitionError = "error"
)

// Client object used to interact with rancher server
type Client struct {
	AuthenticationType ops.AuthenticationType // denotes the authenticate type to contact cluster orchestrator
	rancherAPIClient   apiClient
}