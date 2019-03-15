package cmd

import (
	"github.com/matthieuberger/k8s-cluster/src/generator"
	"github.com/matthieuberger/k8s-cluster/src/parser"
	"github.com/spf13/cobra"
)

func init() {
	createCmd.AddCommand(hostsCmd)
}

var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Generate the hosts.ini file from the terraform.tfsate file",
	Run: func(cmd *cobra.Command, args []string) {
		// Step 1 - Check if terraform.tfstate file exists

		// Steap 2 - Parse terraform.tfstate
		terraformState, err := parser.ParseTerraformStateFile("terraform.tfstate")
		if err != nil {
			panic(err)
		}

		serverList := parser.ServerFromTerraformState(terraformState)
		hosts := parser.ServerListToHostsFile(serverList)
		config := &generator.HostsConfig{Hosts: serverList, Tags: hosts}

		// Step 3 - Write hosts.ini file
		generator.HydrateHosts(config)
	},
}
