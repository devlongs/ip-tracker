package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Iptracker CLI app",
		Long: "My Iptracker CLI application",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}