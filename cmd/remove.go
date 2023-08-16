// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove name [flags]",
	Short: "Remove a clusterset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Printf("Removed clusterset %q\n", args[0])
		}
	},
}

func init() {
	clustersetCmd.AddCommand(removeCmd)
}
