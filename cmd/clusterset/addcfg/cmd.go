// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package addcfg

import (
	"fmt"
	"strings"

	"github.com/nirs/kubectl-ramen/cmd/options"
	"github.com/spf13/cobra"
)

var (
	hubKubeconfig      string
	cluster1Kubeconfig string
	cluster2Kubeconfig string

	description = `
This command creates a clusterset from kubeconfigs files
created from the hub and the managed clusters
`

	Command = &cobra.Command{
		Use:   "add-cfg name [flags]",
		Short: "Add a clusterset from kubeconfigs files",
		Long:  strings.Trim(description, "\n"),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if options.Verbose {
				fmt.Printf("Adding clusterset %q from kubeconfigs\n", args[0])
				fmt.Printf("  hub: %q\n", hubKubeconfig)
				fmt.Printf("  cluster1: %q\n", cluster1Kubeconfig)
				fmt.Printf("  cluster2: %q\n", cluster2Kubeconfig)
			}
		},
	}
)

func init() {
	// Sorting flags messes up the help text.
	Command.Flags().SortFlags = false

	Command.Flags().StringVar(&hubKubeconfig, "hub", "", "hub kubeconfig file")
	Command.MarkFlagRequired("hub")

	Command.Flags().StringVar(&cluster1Kubeconfig, "cluster1", "", "cluster1 kubeconfig file")
	Command.MarkFlagRequired("cluster1")

	Command.Flags().StringVar(&cluster2Kubeconfig, "cluster2", "", "cluster2 kubeconfig file")
	Command.MarkFlagRequired("cluster2")
}
