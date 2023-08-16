// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool

	rootExample = `  # Add clusterset from drenv environment file
  kubectl ramen clusterset add-env rdr --env-file regional-dr.yaml

  # Add clusterset from kubeconfigs
  kubectl ramen clusterset add-cfg rdr --hub hub.cfg --cluster1 c1.cfg --cluster2 c2.cfg

  # Deploy ramen in clusterset "rdr"
  kubectl ramen deploy --clusterset rdr

  # Enable DR for application "busybox-sample"
  kubectl ramen enable --clusterset rdr --namespace busybox-sample

  # Get application "busybox-sample"
  kubectl ramen status --clusterset rdr --namespace busybox-sample

  # Watch application status
  kubectl ramen status --clusterset rdr --namespace busybox-sample --watch

  # Failover application to the second cluster
  kubectl ramen failover --clusterset rdr --namespace busybox-sample

  # Relocate application to the second cluster
  kubectl ramen relocate --clusterset rdr --namespace busybox-sample

  # Disable DR for the busybox-sample application
  kubectl ramen disable --clusterset rdr --namespace busybox-sample

  # Undeploy ramen in clusterset "rdr"
  kubectl ramen undeploy --clusterset rdr --namespace busybox-sample`
)

var rootCmd = &cobra.Command{
	// NOTE: using No-Break-Space (U+00A0) since only the first word
	// is considered as the command name.
	Use:     "kubectl\u00A0ramen",
	Short:   "The kubectl ramen plugin manages Ramen DR",
	Example: rootExample,
}

func init() {
	// Global flags shared by sub commands.
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "be more verbose")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
