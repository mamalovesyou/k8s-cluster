package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of k8s-cluster",
	Long:  `All software has versions. This is k8s-cluster's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("k8s-cluster: ", version)
	},
}
