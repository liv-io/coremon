package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use: "show",
	Aliases: []string{"sho", "sh", "status", "statu", "stat", "sta", "st", "get", "ge"},
	Short: "Shows detailed information about [unit]",
	Long: `
The 'show' command shows detailed information about the called [unit].
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
Usage:
  coremon show [unit]

Examples:
  coremon sh io
  coremon sho cpu
  coremon show disk

Use "coremon show --help" for more information on the available units.
`)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	rootCmd.PersistentFlags().IntP("interval", "i", 1, "Interval in seconds between each report")
}
