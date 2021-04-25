module github.com/adarsh-zededa/cluster-orchestrator

go 1.15

require (
	github.com/rancher/norman v0.0.0-20210423002317-8e6ffc77a819
	github.com/rancher/rancher/pkg/apis v0.0.0-20210425051802-2d7fa0eab3e8
	github.com/rancher/rancher/pkg/client v0.0.0-20210425051802-2d7fa0eab3e8
	github.com/sirupsen/logrus v1.8.1
)

replace k8s.io/client-go => k8s.io/client-go v0.21.0
