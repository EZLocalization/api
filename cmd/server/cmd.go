package main

import (
	"github.com/spf13/cobra"
)

// NewCommandLine Create ComandLine instance
func NewCommandLine() *cobra.Command {
	var configPath string
	var versionType string

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start Admin API Server",
		Long:  `Start Admin API Server`,
		Run: func(cmd *cobra.Command, args []string) {
			StartServer(configPath)
		},
	}
	startCmd.Flags().StringVarP(&configPath, "config", "c", "./config.json", "path of configruation file")

	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			if len(versionType) >= 1 {
				// TDO: Version showing
				return
			}

			cmd.Help()
		},
	}
	rootCmd.PersistentFlags().StringVarP(&versionType, "version", "v", "", "show version (values: text, json)")
	rootCmd.AddCommand(startCmd)

	return rootCmd
}
