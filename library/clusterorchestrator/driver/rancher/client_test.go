package rancher

import (
	"reflect"
	"testing"

	"github.com/adarsh-zededa/cluster-orchestrator/library/clusterorchestrator/ops"
)

func Test_getCPUUsagePercentage(t *testing.T) {
	type args struct {
		usedCPU   string
		totalCPUs int
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "test_getCPUUsagePercentage",
			args: args{
				usedCPU:   "100m",
				totalCPUs: 2,
			},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCPUUsagePercentage(tt.args.usedCPU, tt.args.totalCPUs)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCPUUsagePercentage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCPUUsagePercentage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMemoryInBytes(t *testing.T) {
	type args struct {
		memory string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test_getMemoryInBytes",
			args: args{
				memory: "1000Ki",
			},
			want:    1024000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMemoryInBytes(tt.args.memory)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMemoryInBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getMemoryInBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseClusterSummary(t *testing.T) {
	type args struct {
		clusterSummary *Cluster
		clusterNodes   []*ops.Node
	}
	tests := []struct {
		name    string
		args    args
		want    *ops.ClusterStatus
		wantErr bool
	}{
		{
			name: "test_parseClusterSummary",
			args: args{
				clusterSummary: &Cluster{
					Name: "testClusterName",
					Resource: Resource{
						ID: "testClusterID",
					},
					TransitioningMessage: "testTransitioningMessage",
					State:                "running",
					Allocatable: map[string]string{
						"cpu":    "2",
						"memory": "10000Ki",
						"pods":   "110",
					},
					Requested: map[string]string{
						"cpu":    "100m",
						"memory": "100Ki",
						"pods":   "6",
					},
				},
				clusterNodes: []*ops.Node{
					{
						Name:               "testClusterNodeName",
						ID:                 "testClusterNodeID",
						ClusterID:          "testClusterID",
						NodeIP:             "0.0.0.0",
						ErrorString:        "",
						TotalCPUs:          2,
						TotalMemoryInBytes: 10240000,
						TotalPodsCapacity:  110,
						UsedCPUPercentage:  5,
						UsedMemoryInBytes:  102400,
						UsedPods:           4,
						State:              0,
						Role:               0,
					},
				},
			},
			want: &ops.ClusterStatus{
				Name: "testClusterName",
				ID:   "testClusterID",
				Nodes: []*ops.Node{
					{
						Name:               "testClusterNodeName",
						ID:                 "testClusterNodeID",
						ClusterID:          "testClusterID",
						NodeIP:             "0.0.0.0",
						ErrorString:        "",
						TotalCPUs:          2,
						TotalMemoryInBytes: 10240000,
						TotalPodsCapacity:  110,
						UsedCPUPercentage:  5,
						UsedMemoryInBytes:  102400,
						UsedPods:           4,
						State:              0,
						Role:               0,
					},
				},
				State: ops.STATE_ONLINE,
				Metrics: ops.Metrics{
					CPUPercentage:    5,
					PodsPercentage:   5.4545455,
					MemoryPercentage: 1,
				},
				ErrorString: "testTransitioningMessage",
				NodeErrors:  make(map[string]string),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseClusterSummary(tt.args.clusterSummary, tt.args.clusterNodes)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseClusterSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseClusterSummary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseNodeSummary(t *testing.T) {
	type args struct {
		nodeSummary Node
	}
	tests := []struct {
		name    string
		args    args
		want    *ops.Node
		wantErr bool
	}{
		{
			name: "test_parseNodeSummary",
			args: args{
				nodeSummary: Node{
					NodeName: "testClusterNodeName",
					Resource: Resource{
						ID: "testClusterNodeID",
					},
					ClusterID:            "testClusterID",
					IPAddress:            "0.0.0.0",
					TransitioningMessage: "testTransitioningMessage",
					ControlPlane:         true,
					Worker:               false,
					Allocatable: map[string]string{
						"cpu":    "2",
						"memory": "10000Ki",
						"pods":   "110",
					},
					Requested: map[string]string{
						"cpu":    "100m",
						"memory": "100Ki",
						"pods":   "6",
					},
					State: "provisioned",
				},
			},
			want: &ops.Node{
				Name:               "testClusterNodeName",
				ID:                 "testClusterNodeID",
				ClusterID:          "testClusterID",
				NodeIP:             "0.0.0.0",
				ErrorString:        "testTransitioningMessage",
				TotalCPUs:          2,
				TotalMemoryInBytes: 10240000,
				TotalPodsCapacity:  110,
				UsedCPUPercentage:  5,
				UsedMemoryInBytes:  102400,
				UsedPods:           6,
				State:              ops.STATE_ONLINE,
				Role:               ops.NODE_ROLE_SERVER,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseNodeSummary(tt.args.nodeSummary)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNodeSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNodeSummary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getState(t *testing.T) {
	type args struct {
		rancherState string
	}
	tests := []struct {
		name string
		args args
		want ops.State
	}{
		{
			name: "test_getState",
			args: args{
				rancherState: "running",
			},
			want: ops.STATE_ONLINE,
		},
		{
			name: "test_getState",
			args: args{
				rancherState: "random",
			},
			want: ops.STATE_ERROR,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getState(tt.args.rancherState); got != tt.want {
				t.Errorf("getState() = %v, want %v", got, tt.want)
			}
		})
	}
}
