package ops

// OrchestratorClient defines methods to establish and verify connection with the cluster orchestrator
type OrchestratorClient interface {
	// Login authenticates using UserName and Password
	Login() (string, error)
	// VerifyConnection verifies if authentication API token is valid.
	VerifyTokenValidity() error
	// CreateCluster creates a cluster with the given name and return ClusterConfig object
	CreateCluster(clusterName string) (*ClusterConfig, error)
	// GetClusterStatus reads cluster status by cluster ID and returns ClusterStatus object
	GetClusterStatusByID(clusterID string) (*ClusterStatus, error)
	// ListClusterStatuses returns list of cluster statuses
	ListClusterStatuses() ([]*ClusterStatus, error)
	// DeleteCluster removes a cluster by ID from cluster orchestrator
	DeleteCluster(clusterID string) error
}