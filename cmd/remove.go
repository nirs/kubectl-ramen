// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"log"

	"github.com/nirs/kubectl-ramen/config"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove name [flags]",
	Short: "Remove a clusterset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		store := config.DefaultStore()
		err := store.RemoveClusterSet(args[0])
		if err != nil {
			log.Fatalf("Cannot remove clusterset: %s", err)
		}
	},
}

func init() {
	clustersetCmd.AddCommand(removeCmd)
}
