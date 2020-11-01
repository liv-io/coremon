package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ioCmd = &cobra.Command{
	Use: "io",
	Aliases: []string{"iostat", "iosta", "iost", "ios", "iowait", "iowai", "iowa", "iow"},
	Short: "Shows detailed IO information",
	Long: `
The 'io' command shows detailed information about the Input/Output.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`UNIT		TYPE		STATUS

IOWait		system		[|                                         1.4%]
`)
	},
}

func init() {
	showCmd.AddCommand(ioCmd)
}
