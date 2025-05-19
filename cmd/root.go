// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of zcp
package cmd

import (
	"os"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/zachrundle/zcp/internal/models"
	"github.com/zachrundle/zcp/internal/setup"
)

const (
	appName      = "zcp"
	shortAppDesc = "A tui application to assist building install-configs.yaml files for Openshift clusters"
	longAppDesc  = `A tui application that uses API calls to assist in creating the install-config.yaml files
for build Openshift clusters. By dynamically pulling in values, it makes creating configs less prone to human error.

AWS module: 	github.com/aws/aws-sdk-go-v2
vSphere module: github.com/vmware/govmomi`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: shortAppDesc,
	Long: longAppDesc,
 	RunE: func(cmd *cobra.Command, args []string) error {
		launchTUI()
	// RunE: func(cmd *cobra.Command, args []string) error {
		// p := tea.NewProgram(model.InitialModel())
		// _, err := p.Run()
		// return err
	return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("file", "f", false, "Load in an existing install-config.yaml")
}

func launchTUI() {
	if !setup.BaseTemplateExists() {
		cfg := setup.New()
		model := model.InitialModel(cfg)
		p := tea.NewProgram(model, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error in setup: %v\n", err)
			os.Exit(1)
		}
	}
}
