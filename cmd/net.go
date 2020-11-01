package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var netCmd = &cobra.Command{
	Use: "net",
	Aliases: []string{"networking", "networkin", "networki", "network", "networ", "netwo", "netw", "ne"},
	Short: "Shows detailed network information",
	Long: `
The 'net' command shows detailed information about the network utilization.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`UNIT		TYPE		STATUS

eth0: in	network		[|                                      9.6Mbps]
eth0: out	network		[                                       0.5Mbps]
lo: in		network		[                                       0.0Mbps]
lo: out		network		[                                       0.0Mbps]
`)
	},
}

func init() {
	showCmd.AddCommand(netCmd)
}
