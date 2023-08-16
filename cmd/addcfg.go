// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	hubKubeconfig      string
	cluster1Kubeconfig string
	cluster2Kubeconfig string
)

var addCfgCmd = &cobra.Command{
	Use:   "add-cfg name [flags]",
	Short: "Add a clusterset from kubeconfigs files",
	Long: `This command creates a clusterset from kubeconfigs files
created from the hub and the managed clusters`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Printf("Adding clusterset %q from kubeconfigs\n", args[0])
			fmt.Printf("  hub: %q\n", hubKubeconfig)
			fmt.Printf("  cluster1: %q\n", cluster1Kubeconfig)
			fmt.Printf("  cluster2: %q\n", cluster2Kubeconfig)
		}
	},
}

func init() {
	// Sorting flags messes up the help text.
	addCfgCmd.Flags().SortFlags = false

	addCfgCmd.Flags().StringVar(&hubKubeconfig, "hub", "", "hub kubeconfig file")
	addCfgCmd.MarkFlagRequired("hub")

	addCfgCmd.Flags().StringVar(&cluster1Kubeconfig, "cluster1", "", "cluster1 kubeconfig file")
	addCfgCmd.MarkFlagRequired("cluster1")

	addCfgCmd.Flags().StringVar(&cluster2Kubeconfig, "cluster2", "", "cluster2 kubeconfig file")
	addCfgCmd.MarkFlagRequired("cluster2")

	clustersetCmd.AddCommand(addCfgCmd)
}
