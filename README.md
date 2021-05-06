# cluster-orchestrator
A GO client library to interact with cluster orchestrators (ex. Rancher)

### APIs exposed in the orchestrator library
____________
```
Login() error
VerifyTokenValidity() error
CreateCluster(clusterName string) (*ClusterConfig, error)
GetClusterStatusByID(clusterID string) (*ClusterStatus, error)
ListClusterStatuses() ([]*ClusterStatus, error)
DeleteCluster(clusterID string) error
```
Refer `./testclient.go` on how to create and interact with the orchestrator client

