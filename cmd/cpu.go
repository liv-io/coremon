package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use: "cpu",
	Aliases: []string{"cp", "processor", "processo", "process", "proces", "proce", "proc", "pro", "pr"},
	Short: "Shows detailed CPU information",
	Long: `
The 'cpu' command shows detailed information about the processor utilization.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`UNIT		TYPE		STATUS

Processor	system		[||||                                     10.0%]
 1		system		[|                                         2.7%]
 2		system		[|                                         1.2%]
 3		system		[||                                        4.7%]
 4		system		[|                                         1.4%]
`)
	},
}

func init() {
	showCmd.AddCommand(cpuCmd)
}
