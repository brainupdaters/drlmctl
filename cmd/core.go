// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import "github.com/spf13/cobra"

var coreCmd = &cobra.Command{
	Use:   "core",
	Short: "TODO",
	Long:  `TODO`,
}

func init() {
	rootCmd.AddCommand(coreCmd)
}
