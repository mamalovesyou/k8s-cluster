package cmd

import (
	"github.com/matthieuberger/k8s-cluster/src/generator"
	"github.com/spf13/cobra"
)

var Destination string
var ClusterName string
var Region string

var NumberOfLb int
var NumberOfEtcd int
var NumberOfMasters int
var NumberOfWorkers int
var SeparateEtcd bool

var LBInstanceType string
var EtcdInstanceType string
var MasterInstanceType string
var WorkerInstanceType string

func init() {

	generateCmd.Flags().StringVarP(&Destination, "output", "o", "./", "Output directory.")
	generateCmd.Flags().StringVarP(&ClusterName, "name", "n", "mycluster", "Name of the cluster")
	generateCmd.Flags().StringVarP(&Region, "region", "r", "par1", "The provider region")

	generateCmd.Flags().IntVarP(&NumberOfLb, "load-balancer", "", 2, "Number of load balancer nodes")
	generateCmd.Flags().IntVarP(&NumberOfEtcd, "etcd", "", 3, "Number of etcd nodes")
	generateCmd.Flags().IntVarP(&NumberOfMasters, "kube-master", "", 2, "Number of kubernetes master nodes")
	generateCmd.Flags().IntVarP(&NumberOfWorkers, "kube-worker", "", 2, "Number of kubernetes workers nodes")
	generateCmd.Flags().BoolVarP(&SeparateEtcd, "seprate-etcd", "", false, "If true it will use load balancers as etcd as well")

	generateCmd.Flags().StringVarP(&LBInstanceType, "lb-type", "", "Start-S", "Instance type used for load balancers nodes")
	generateCmd.Flags().StringVarP(&EtcdInstanceType, "etcd-type", "", "Start-S", "Instance type used for etcds nodes")
	generateCmd.Flags().StringVarP(&MasterInstanceType, "master-type", "", "Start-S", "Instance type used for masters nodes")
	generateCmd.Flags().StringVarP(&WorkerInstanceType, "worker-type", "", "Start-M", "Instance type used for workers nodes")

	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate all needed files to setup a cluster with terraform",
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.CreateAllDirs(Destination, ClusterName)
		if err != nil {
			panic(err)
		}

		clusterConfig := &generator.ClusterConfig{
			Region:             Region,
			ClusterName:        ClusterName,
			LbInstanceType:     LBInstanceType,
			EtcdInstanceType:   EtcdInstanceType,
			MasterInstanceType: MasterInstanceType,
			WorkerInstanceType: WorkerInstanceType,
			UseLbAsEtcd:        !SeparateEtcd,
			LoadBalancerNodes:  NumberOfLb,
			EtcdNodes:          NumberOfEtcd,
			KubeMasterNodes:    NumberOfMasters,
			KubeWorkerNodes:    NumberOfWorkers,
		}
		terraformConfig := generator.NewTerraformConfig(clusterConfig, Destination)
		generator.HydrateTerraformCluster(terraformConfig)
		generator.HydrateTerraformVariables(terraformConfig)
	},
}
