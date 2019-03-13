package generator

import "path"

// the instance struct
type Instance struct {
	Type  string
	State string
	Tags  []string
}

// The structure of the cluster config
type ClusterConfig struct {
	ClusterName string
	Region      string

	LoadBalancerNodes int
	KubeMasterNodes   int
	KubeWorkerNodes   int
	EtcdNodes         int

	UseLbAsEtcd bool

	LbInstanceType     string
	EtcdInstanceType   string
	MasterInstanceType string
	WorkerInstanceType string
}

// The structure of the cluster config
type TerraformConfig struct {
	DestinationPath string

	ClusterName   string
	Region        string
	NumberOfNodes int

	LoadBalancerNodes []*Instance
	KubeMasterNodes   []*Instance
	KubeWorkerNodes   []*Instance
	EtcdNodes         []*Instance

	IpCounter   int
	NodeCounter int
}

func NewTerraformConfig(clusterConfig *ClusterConfig, path string) *TerraformConfig {

	if clusterConfig.UseLbAsEtcd {
		realEtcdNodes := clusterConfig.EtcdNodes - clusterConfig.LoadBalancerNodes
		if realEtcdNodes > 0 {
			clusterConfig.EtcdNodes = realEtcdNodes
		} else {
			clusterConfig.EtcdNodes = 0
		}
	}

	nodesCount := 0

	// Create Load Balancers Instance
	lbInstances := make([]*Instance, clusterConfig.LoadBalancerNodes)
	for i := 0; i < clusterConfig.LoadBalancerNodes; i++ {
		lbInstances[i] = &Instance{Type: clusterConfig.LbInstanceType, State: "running"}
		tags := []string{"load-balancer", clusterConfig.ClusterName}
		if clusterConfig.UseLbAsEtcd {
			tags = append(tags, "etcd")
		}
		lbInstances[i].Tags = tags
		nodesCount += 1
	}

	// Create Kube Masters Instance
	masterInstances := make([]*Instance, clusterConfig.KubeMasterNodes)
	for i := 0; i < clusterConfig.KubeMasterNodes; i++ {
		masterInstances[i] = &Instance{Type: clusterConfig.MasterInstanceType, State: "running"}
		tags := []string{"kube-master", clusterConfig.ClusterName}
		masterInstances[i].Tags = tags
		nodesCount += 1
	}

	// Create Kube Workers Instance
	workerInstances := make([]*Instance, clusterConfig.KubeWorkerNodes)
	for i := 0; i < clusterConfig.KubeWorkerNodes; i++ {
		workerInstances[i] = &Instance{Type: clusterConfig.WorkerInstanceType, State: "running"}
		tags := []string{"kube-worker", clusterConfig.ClusterName}
		workerInstances[i].Tags = tags
		nodesCount += 1
	}

	// Create Etcd Instance
	etcdInstances := make([]*Instance, clusterConfig.EtcdNodes)
	for i := 0; i < clusterConfig.EtcdNodes; i++ {
		etcdInstances[i] = &Instance{Type: clusterConfig.EtcdInstanceType, State: "running"}
		tags := []string{"etcd", clusterConfig.ClusterName}
		etcdInstances[i].Tags = tags
		nodesCount += 1
	}

	return &TerraformConfig{
		DestinationPath:   path,
		NumberOfNodes:     nodesCount,
		ClusterName:       clusterConfig.ClusterName,
		Region:            clusterConfig.Region,
		LoadBalancerNodes: lbInstances,
		KubeMasterNodes:   masterInstances,
		KubeWorkerNodes:   workerInstances,
		EtcdNodes:         etcdInstances,
		IpCounter:         0,
		NodeCounter:       1,
	}
}

// Counter for the template
func (c *TerraformConfig) IncrementIpCounter() int {
	c.IpCounter += 1
	return c.IpCounter
}

// Counter for the template
func (c *TerraformConfig) IncrementNodeCounter() int {
	c.NodeCounter += 1
	return c.NodeCounter
}

func (c *TerraformConfig) ClusterFilePath() string {
	return path.Join(c.DestinationPath, c.ClusterName, "terraform", c.ClusterName+".tf")
}
