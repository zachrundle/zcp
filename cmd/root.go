// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of zcp
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zcp",
	Short: "A tui application to assist building install-configs.yaml files for Openshift clusters",
	Long: `A tui application that uses API calls to assist in creating the install-config.yaml files
for build Openshift clusters. By dynamically pulling in values, it makes creating configs less prone to human error.

AWS module: 	github.com/aws/aws-sdk-go-v2
vSphere module: github.com/vmware/govmomi`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zcp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("file", "f", false, "Load in an existing install-config.yaml")
}
