// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package remove

import (
	"fmt"

	"github.com/nirs/kubectl-ramen/cmd/options"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "remove name [flags]",
		Short: "Remove a clusterset",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if options.Verbose {
				fmt.Printf("Removed clusterset %q\n", args[0])
			}
		},
	}
)
