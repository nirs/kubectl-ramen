// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"log"

	"github.com/nirs/kubectl-ramen/config"
	"github.com/nirs/kubectl-ramen/config/kube"
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
		configset, err := kube.NewConfigSet(hubKubeconfig, cluster1Kubeconfig, cluster2Kubeconfig)
		if err != nil {
			log.Fatalf("Canot load configs: %s", err)
		}
		store := config.DefaultStore()
		err = store.AddClusterSetFromConfigs(args[0], configset)
		if err != nil {
			log.Fatalf("Cannot add clusterset: %s", err)
		}
	},
}

func init() {
	// Sorting flags messes up the help text.
	addCfgCmd.Flags().SortFlags = false

	addCfgCmd.Flags().StringVar(&hubKubeconfig, "hub", "", "hub kubeconfig file")
	if err := addCfgCmd.MarkFlagRequired("hub"); err != nil {
		log.Fatal(err)
	}

	addCfgCmd.Flags().StringVar(&cluster1Kubeconfig, "cluster1", "", "cluster1 kubeconfig file")
	if err := addCfgCmd.MarkFlagRequired("cluster1"); err != nil {
		log.Fatal(err)
	}

	addCfgCmd.Flags().StringVar(&cluster2Kubeconfig, "cluster2", "", "cluster2 kubeconfig file")
	if err := addCfgCmd.MarkFlagRequired("cluster2"); err != nil {
		log.Fatal(err)
	}

	clustersetCmd.AddCommand(addCfgCmd)
}
