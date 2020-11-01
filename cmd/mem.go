package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memCmd = &cobra.Command{
	Use: "mem",
	Aliases: []string{"memory", "memor", "memo", "me", "swap", "swa", "sw"},
	Short: "Shows detailed memory information",
	Long: `
The 'mem' command shows detailed information about the memory usage and utilization.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`UNIT		TYPE		STATUS

Memory		system		[|||||||||||||||||||                 2.71G/8.0G]
Swap		system		[|                                  16.9M/2.00G]
`)
	},
}

func init() {
	showCmd.AddCommand(memCmd)
}
