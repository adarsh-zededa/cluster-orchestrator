package clusterorchestrator

import (
	"fmt"

	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/driver/rancher"
	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
)

func NewClient(orchestratorConfig ops.OrchestratorConfig) (ops.OrchestratorClient, error) {
	switch orchestratorConfig.Type {
	case ops.ORCHESTRATOR_TYPE_RANCHER:
		if orchestratorConfig.Rancher == nil {
			return nil, fmt.Errorf("rancher config missing")
		}
		rancherClient := &rancher.Client{
			Server:             orchestratorConfig.Rancher.Server,
			Port:               orchestratorConfig.Rancher.Port,
			AuthenticationType: orchestratorConfig.Rancher.AuthenticationType,
			UserName:           orchestratorConfig.Rancher.UserName,
			Password:           orchestratorConfig.Rancher.Password,
			APIToken:           orchestratorConfig.Rancher.APIToken,
		}
		return rancherClient, nil
	default:
		return nil, fmt.Errorf("unrecognized orchestrator type: %d", orchestratorConfig.Type)
	}
}
