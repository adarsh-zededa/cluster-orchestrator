package clusterorchestrator

import (
	"fmt"

	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/driver/dummy"
	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/driver/rancher"
	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
)

func NewClient(orchestratorConfig ops.OrchestratorConfig) (ops.OrchestratorClient, error) {
	switch orchestratorConfig.Type {
	case ops.ORCHESTRATOR_TYPE_RANCHER:
		if orchestratorConfig.Rancher == nil {
			return nil, fmt.Errorf("rancher config missing")
		}
		rancherClient := rancher.NewRancherClient(orchestratorConfig.Rancher)
		return rancherClient, nil
	case ops.ORCHESTRATOR_TYPE_DUMMY:
		if orchestratorConfig.Dummy == nil {
			return nil, fmt.Errorf("dummy config missing")
		}
		dummyClient := dummy.NewDummyClient(orchestratorConfig.Dummy)
		return dummyClient, nil
	default:
		return nil, fmt.Errorf("unrecognized orchestrator type: %d", orchestratorConfig.Type)
	}
}