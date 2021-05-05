package dummy

import (
	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
)

// Client object used to interact with dummy server
type Client struct {
	AuthenticationType ops.AuthenticationType // denotes the authenticate type to contact cluster orchestrator=
}
