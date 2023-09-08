// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"log"

	"github.com/nirs/kubectl-ramen/config"
	"github.com/nirs/kubectl-ramen/config/drenv"
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
		env, err := drenv.Load(envFile, drenv.Options{NamePrefix: namePrefix})
		if err != nil {
			log.Fatalf("Cannot load env file: %s", err)
		}
		store := config.DefaultStore()
		err = store.AddClusterSetFromEnv(args[0], env)
		if err != nil {
			log.Fatalf("Cannot add clusterset: %s", err)
		}
	},
}

func init() {
	// Sorting flags messes up the help text.
	addEnvCmd.Flags().SortFlags = false

	addEnvCmd.Flags().StringVarP(&envFile, "env-file", "f", "", "drenv environment file")
	if err := addEnvCmd.MarkFlagRequired("env-file"); err != nil {
		log.Fatal(err)
	}

	addEnvCmd.Flags().StringVar(&namePrefix, "name-prefix", "", "prefix cluster names")

	clustersetCmd.AddCommand(addEnvCmd)
}
