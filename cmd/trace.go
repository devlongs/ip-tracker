package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "trace the ip",
	Long: `trace the ip address`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testing trace the ip")
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
