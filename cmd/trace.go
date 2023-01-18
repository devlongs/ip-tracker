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
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
			}
		} else {
			fmt.Println("Please provide an IP address to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
