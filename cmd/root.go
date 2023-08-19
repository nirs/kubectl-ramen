// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	verbose bool

	rootExample = `  # Add clusterset from drenv environment file
  %[1]s clusterset add-env rdr --env-file regional-dr.yaml

  # Add clusterset from kubeconfigs
  %[1]s clusterset add-cfg rdr --hub hub.cfg --cluster1 c1.cfg --cluster2 c2.cfg

  # Deploy ramen in clusterset "rdr"
  %[1]s deploy --clusterset rdr

  # Enable DR for application "busybox-sample"
  %[1]s enable --clusterset rdr --namespace busybox-sample

  # Print application "busybox-sample" status
  %[1]s status --clusterset rdr --namespace busybox-sample

  # Watch application "busybox-sample" status
  %[1]s status --clusterset rdr --namespace busybox-sample --watch

  # Failover application to the other managed cluster
  %[1]s failover --clusterset rdr --namespace busybox-sample

  # Relocate application to the other managed cluster
  %[1]s relocate --clusterset rdr --namespace busybox-sample

  # Disable DR for the busybox-sample application
  %[1]s disable --clusterset rdr --namespace busybox-sample

  # Undeploy ramen in clusterset "rdr"
  %[1]s undeploy --clusterset rdr`
)

// Can be `kubectl-ramen` or `oc-ramen`.
var executableName = path.Base(os.Args[0])

var rootCmd = &cobra.Command{
	// We cannot use `kubectl ramen` since the only first word is used by
	// cobra. See https://github.com/spf13/cobra/issues/2017.
	Use:     executableName,
	Short:   fmt.Sprintf("The %[1]s plugin manages Ramen DR", executableName),
	Example: fmt.Sprintf(rootExample, executableName),
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
