// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of zcp
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// platformCmd represents the platform command
var platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Config requirements for supported platforms",
	Long: ` The platform subcommand lists all the supported platforms and their config requirements
that can be used to generate an install-config.yaml for Openshift clusters.
For example:

zts platform
zts platform show`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("List of supported platforms:\n - %s\n - %s\n", "aws", "vsphere")
	},
}

func init() {
	rootCmd.AddCommand(platformCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// platformCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// platformCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
