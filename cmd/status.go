// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/nirs/kubectl-ramen/apps"
	"github.com/nirs/kubectl-ramen/config"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

type output struct {
	Time   time.Time               `json:"time"`
	Status *apps.ApplicationStatus `json:"status"`
}

var clusterset string
var namespace string

var statusCmd = &cobra.Command{
	Use:   "status [flags]",
	Short: "Print application status on a clusterset",
	Run: func(cmd *cobra.Command, args []string) {
		store := config.DefaultStore()
		cs, err := store.GetClusterSet(clusterset)
		if err != nil {
			log.Fatalf("Cannot locate clusterset: %s", err)
		}
		watcher, err := apps.NewApplicationWatcher(cs, namespace)
		if err != nil {
			log.Fatalf("Cannot watch application: %s", err)
		}
		status, err := watcher.Status()
		if err != nil {
			log.Fatalf("Cannot get status: %s", err)
		}
		printStatus(status)
	},
}

func init() {
	statusCmd.Flags().StringVarP(&clusterset, "clusterset", "c", "", "clusterset name")
	addEnvCmd.MarkFlagRequired("clusterset")

	statusCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "application namespace")
	addEnvCmd.MarkFlagRequired("namespace")

	rootCmd.AddCommand(statusCmd)
}

func printStatus(status *apps.ApplicationStatus) {
	out, err := yaml.Marshal(output{
		Time:   time.Now(),
		Status: status,
	})
	if err != nil {
		log.Fatalf("Cannot format status: %s", err)
	}

	fmt.Println("---")
	fmt.Print(string(out))
}
