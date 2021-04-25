package main

import (
	"time"

	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator"
	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
	"github.com/sirupsen/logrus"
)

func main() {
	orchestratorConfig := ops.OrchestratorConfig{
		Type: ops.ORCHESTRATOR_TYPE_RANCHER,
		Rancher: &ops.RancherOrchestratorConfig{
			Server:             "<server-address>",
			Port:               "<server-port>",
			AuthenticationType: ops.AUTHENTICATION_CREDENTIAL,
			UserName:           "<username>",
			Password:           "<password>",
		},
	}
	rancherClient, err := clusterorchestrator.NewClient(orchestratorConfig)
	if err != nil {
		logrus.Errorf("exception while initializing Rancher Client. %v", err)
	}
	startTime := time.Now()
	if err := rancherClient.Login(); err != nil {
		logrus.Errorf("exception while authenticaing Rancher Client. %v", err)
	}

	if err := rancherClient.VerifyTokenValidity(); err != nil {
		logrus.Errorf("exception while verifying Rancher Client token. %v", err)
	}

	clusterName := "adarsh-test-cluster"
	var clusterConfig *ops.ClusterConfig
	if clusterConfig, err = rancherClient.CreateCluster(clusterName); err != nil {
		logrus.Errorf("exception while verifying creating cluster. %v", err)
	}

	if _, err := rancherClient.GetClusterStatus(clusterConfig.ID); err != nil {
		logrus.Errorf("exception while fetching Rancher Cluster status. %v", err)
	}

	if err := rancherClient.DeleteCluster(clusterConfig.ID); err != nil {
		logrus.Errorf("exception while deleting Rancher Cluster status. %v", err)
	}
	duration := time.Now().Sub(startTime)
	logrus.Infof("Overall time take: %v", duration)
}
