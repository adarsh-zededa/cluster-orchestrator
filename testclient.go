package main

import (
	"time"

	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator"
	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
	log "github.com/sirupsen/logrus"
)

func main() {
	orchestratorConfig := ops.OrchestratorConfig{
		Type: ops.ORCHESTRATOR_TYPE_RANCHER,
		Rancher: &ops.RancherOrchestratorConfig{
			Server:             "40.114.78.211",
			Port:               "443",
			AuthenticationType: ops.AUTHENTICATION_CREDENTIAL,
			UserName:           "admin",
			Password:           "k3sAdmin@1234",
		},
	}
	rancherClient, err := clusterorchestrator.NewClient(orchestratorConfig)
	if err != nil {
		log.Errorf("exception while initializing Rancher Client. %v", err)
	}
	startTime := time.Now()
	if err := rancherClient.Login(); err != nil {
		log.Errorf("exception while authenticaing Rancher Client. %v", err)
	}

	if err := rancherClient.VerifyTokenValidity(); err != nil {
		log.Errorf("exception while verifying Rancher Client token. %v", err)
	}

	clusterName := "adarsh-test-cluster"
	var clusterConfig *ops.ClusterConfig
	if clusterConfig, err = rancherClient.CreateCluster(clusterName); err != nil {
		log.Errorf("exception while verifying creating cluster. %v", err)
	}

	if _, err := rancherClient.GetClusterStatusByID("c-khqsn"); err != nil {
		log.Errorf("exception while fetching Rancher Cluster status. %v", err)
	}

	if err := rancherClient.DeleteCluster(clusterConfig.ID); err != nil {
		log.Errorf("exception while deleting Rancher Cluster status. %v", err)
	}
	duration := time.Now().Sub(startTime)
	log.Infof("Overall time take: %v", duration)
}
