// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/nirs/kubectl-ramen/config"
	"github.com/nirs/kubectl-ramen/config/envfile"
	"github.com/spf13/cobra"
)

var envFile string
var namePrefix string

var addEnvCmd = &cobra.Command{
	Use:   "add-env name [flags]",
	Short: "Add a clusterset from drenv environment file",
	Long: `This command creates a clusterset from drenv environment
file (mostly useful for Ramen developers)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		env, err := envfile.Load(envFile, envfile.Options{NamePrefix: namePrefix})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot load env file: %s\n", err)
			os.Exit(1)
		}
		store := config.DefaultStore()
		err = store.AddClusterSetFromEnvFile(args[0], env)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot add clusterset: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	// Sorting flags messes up the help text.
	addEnvCmd.Flags().SortFlags = false

	addEnvCmd.Flags().StringVarP(&envFile, "env-file", "f", "", "drenv environment file")
	addEnvCmd.MarkFlagRequired("env-file")

	addEnvCmd.Flags().StringVar(&namePrefix, "name-prefix", "", "prefix cluster names")

	clustersetCmd.AddCommand(addEnvCmd)
}
