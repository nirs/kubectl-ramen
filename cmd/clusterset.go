// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clustersetCmd = &cobra.Command{
	Use:   "clusterset [flags]",
	Short: "Managed clustersets configurations",
	Long: `This commands adds or removes clusterset configurations
that can be used later with the --clusterset option.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fake-clusterset")
	},
}

func init() {
	rootCmd.AddCommand(clustersetCmd)
}
