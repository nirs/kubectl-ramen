// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"log"

	"github.com/nirs/kubectl-ramen/config"
	"github.com/spf13/cobra"
)

var clustersetCmd = &cobra.Command{
	Use:   "clusterset [flags]",
	Short: "Managed clustersets configurations",
	Long: `This commands adds or removes clusterset configurations
that can be used later with the --clusterset option.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		store := config.DefaultStore()

		clustersets, err := store.ListClusterSets()
		if err != nil {
			log.Fatalf("Cannot list clustersets: %s", err)
		}

		for _, name := range clustersets {
			fmt.Println(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(clustersetCmd)
}
