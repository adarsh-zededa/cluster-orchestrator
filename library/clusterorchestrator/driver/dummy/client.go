package dummy

import (
	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
)

func NewDummyClient(dummyConfig *ops.DummyOrchestratorConfig) ops.OrchestratorClient {
	dummyClient := &Client{
		AuthenticationType: dummyConfig.AuthenticationType,
	}
	return dummyClient
}

func (rc *Client) Login() (string, error) {

	return "DummyAPIToken", nil
}

func (rc *Client) VerifyTokenValidity() error {
	return nil
}

func (rc *Client) CreateCluster(clusterName string) (*ops.ClusterConfig, error) {
	clusterConfig := &ops.ClusterConfig{
		Name:        clusterName,
		ID:          clusterName + "ID",
		ManifestURL: "dummyClusterManifest.url",
	}
	return clusterConfig, nil
}

func (rc *Client) GetClusterStatusByID(clusterID string) (*ops.ClusterStatus, error) {

	clusterStatus := &ops.ClusterStatus{
		Name: "dummyClusterName",
		ID:   clusterID,
		Nodes: []*ops.Node{
			{
				Name:               "dummyClusterNodeName",
				ID:                 "dummyClusterNodeID",
				ClusterID:          clusterID,
				NodeIP:             "0.0.0.0",
				ErrorString:        "",
				TotalCPUs:          0,
				TotalMemoryInBytes: 0,
				TotalPodsCapacity:  0,
				UsedCPUPercentage:  0,
				UsedMemoryInBytes:  0,
				UsedPods:           0,
				State:              0,
				Role:               0,
			},
		},
		State: ops.STATE_INIT,
		Metrics: ops.Metrics{
			CPUPercentage:    1,
			PodsPercentage:   1,
			MemoryPercentage: 1,
		},
		ErrorString: "",
		NodeErrors:  nil,
	}
	return clusterStatus, nil
}

func (rc *Client) ListClusterStatuses() ([]*ops.ClusterStatus, error) {
	clusterStatusList := []*ops.ClusterStatus{
		{
			Name: "dummyClusterName",
			ID:   "dummyClusterID",
			Nodes: []*ops.Node{
				{
					Name:               "dummyClusterNodeName",
					ID:                 "dummyClusterNodeID",
					ClusterID:          "dummyClusterID",
					NodeIP:             "0.0.0.0",
					ErrorString:        "",
					TotalCPUs:          0,
					TotalMemoryInBytes: 0,
					TotalPodsCapacity:  0,
					UsedCPUPercentage:  0,
					UsedMemoryInBytes:  0,
					UsedPods:           0,
					State:              0,
					Role:               0,
				},
			},
			State: ops.STATE_INIT,
			Metrics: ops.Metrics{
				CPUPercentage:    1,
				PodsPercentage:   1,
				MemoryPercentage: 1,
			},
			ErrorString: "",
			NodeErrors:  nil,
		},
	}
	return clusterStatusList, nil
}

func (rc *Client) DeleteCluster(clusterID string) error {
	return nil
}
