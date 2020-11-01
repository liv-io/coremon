package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use: "disk",
	Aliases: []string{"dis", "di", "storage", "storag", "stora", "stor", "sto", "st"},
	Short: "Shows detailed disk information",
	Long: `
The 'disk' command shows detailed information about the disk usage and utilization.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`UNIT		TYPE		STATUS

sda1		storage		[|||||||                                  17.2%]
sda2		storage		[||||||||||||                             26.1%]
sdb1		storage		[|||||||||||||||||||||||                  50.6%]
`)
	},
}

func init() {
	showCmd.AddCommand(diskCmd)
}
