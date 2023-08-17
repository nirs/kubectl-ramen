// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/nirs/kubectl-ramen/core"
	"github.com/spf13/cobra"
)

var clustersetCmd = &cobra.Command{
	Use:   "clusterset [flags]",
	Short: "Managed clustersets configurations",
	Long: `This commands adds or removes clusterset configurations
that can be used later with the --clusterset option.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		store := core.DefaultConfigStorage()

		clustersets, err := store.ListClusterSets()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot list clustersets: %s\n", err)
			os.Exit(1)
		}

		for _, name := range clustersets {
			fmt.Println(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(clustersetCmd)
}
