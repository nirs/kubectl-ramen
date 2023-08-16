// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package addenv

import (
	"fmt"
	"strings"

	"github.com/nirs/kubectl-ramen/cmd/options"
	"github.com/spf13/cobra"
)

var (
	envfile string

	description = `
This command creates a clusterset from drenv environment
file (mostly useful for Ramen developers)
`

	Command = &cobra.Command{
		Use:   "add-env name [flags]",
		Short: "Add a clusterset from drenv environment file",
		Long:  strings.Trim(description, "\n"),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if options.Verbose {
				fmt.Printf("Adding clusterset %q from %q\n", args[0], envfile)
			}
		},
	}
)

func init() {
	Command.Flags().StringVarP(
		&envfile, "envfile", "f", "", "ramen drenv environment file")
	Command.MarkFlagRequired("envfile")
}
