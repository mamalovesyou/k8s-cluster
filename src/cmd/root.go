package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "k8s-cluster",
	Short: `k8s-cluster is a very fast config file generator
	to deploy a kubernetes cluster with external load balancer`,
	Long: `Configiguration file generator to deploy a kubernetes cluster
	with external load balancers using Scaleway cloud provider. It generates terraform
	files, ansible playbooks, kubernetes deployements such as dashboard, jenkins or
	ELK stack. The idea is to deploy a custom k8s-cluster for small team or alone dev
	in couples of minutes.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
