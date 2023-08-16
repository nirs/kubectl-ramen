// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package clusterset

import (
	"fmt"

	"github.com/nirs/kubectl-ramen/cmd/clusterset/addcfg"
	"github.com/nirs/kubectl-ramen/cmd/clusterset/addenv"
	"github.com/nirs/kubectl-ramen/cmd/clusterset/remove"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "clusterset [flags]",
		Short: "Managed clustersets",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("fake-clusterset")
		},
	}
)

func init() {
	// Sub commands.
	Command.AddCommand(addenv.Command)
	Command.AddCommand(addcfg.Command)
	Command.AddCommand(remove.Command)
}
