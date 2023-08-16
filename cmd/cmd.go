// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/nirs/kubectl-ramen/cmd/clusterset"
	"github.com/nirs/kubectl-ramen/cmd/options"
	"github.com/spf13/cobra"
)

var (
	example = `
  # Add clusterset from drenv environment file
  %[1]s clusterset add-env rdr --envfile regional-dr.yaml

  # Add clusterset from kubeconfigs
  %[1]s clusterset add-cfg rdr --hub hub.cfg --cluster1 c1.cfg --cluster2 c2.cfg

  # Deploy ramen in clusterset "rdr"
  %[1]s deploy --clusterset rdr

  # Enable DR for application "busybox-sample"
  %[1]s enable --clusterset rdr --namespace busybox-sample

  # Get application "busybox-sample"
  %[1]s status --clusterset rdr --namespace busybox-sample

  # Watch application status
  %[1]s status --clusterset rdr --namespace busybox-sample --watch

  # Failover application to the second cluster
  %[1]s failover --clusterset rdr --namespace busybox-sample

  # Relocate application to the second cluster
  %[1]s relocate --clusterset rdr --namespace busybox-sample

  # Disable DR for the busybox-sample application
  %[1]s disable --clusterset rdr --namespace busybox-sample

  # Undeploy ramen in clusterset "rdr"
  %[1]s undeploy --clusterset rdr --namespace busybox-sample
`
)

var (
	rootCommand = &cobra.Command{
		// NOTE: using No-Break-Space (U+00A0) since only the first word
		// is considered as the command name.
		Use:     "kubectl\u00A0ramen",
		Short:   "The kubectl ramen plugin manages Ramen DR",
		Example: strings.Trim(fmt.Sprintf(example, "kubectl ramen"), "\n"),
	}
)

func init() {
	// Sub commands.
	rootCommand.AddCommand(clusterset.Command)

	// Global flags shared by sub commands.
	rootCommand.PersistentFlags().BoolVarP(
		&options.Verbose, "verbose", "v", false, "be more verbose")
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
