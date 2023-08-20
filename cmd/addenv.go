// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/nirs/kubectl-ramen/config"
	"github.com/spf13/cobra"
)

var envFile string

var addEnvCmd = &cobra.Command{
	Use:   "add-env name [flags]",
	Short: "Add a clusterset from drenv environment file",
	Long: `This command creates a clusterset from drenv environment
file (mostly useful for Ramen developers)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		store := config.DefaultStore()
		err := store.AddClusterSetFromEnvFile(args[0], envFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot add clusterset: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	addEnvCmd.Flags().StringVarP(&envFile, "env-file", "f", "", "drenv environment file")
	addEnvCmd.MarkFlagRequired("env-file")

	clustersetCmd.AddCommand(addEnvCmd)
}
