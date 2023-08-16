// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

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
		if verbose {
			fmt.Printf("Adding clusterset %q from %q\n", args[0], envFile)
		}
	},
}

func init() {
	addEnvCmd.Flags().StringVarP(&envFile, "env-file", "f", "", "drenv environment file")
	addEnvCmd.MarkFlagRequired("env-file")

	clustersetCmd.AddCommand(addEnvCmd)
}
