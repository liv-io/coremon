package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use: "info",
	Aliases: []string{"inf", "in"},
	Short: "Shows general system information",
	Long: `
The 'info' command shows general information about the system.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`KEY			VALUE

Hostname		host01
FQDN			host01.example.com
IP Address		10.1.1.11
Network			10.1.1.0/24
Default Gateway		10.1.1.1

System Start Time	2020-07-01 08:34:02 UTC
System Uptime		11 days 16 hours 38 minutes

OS Name			CentOS
OS Version		8.2.2004
OS Architecture		x86_64

Processor Sockets	1
Processor Cores		4

Memory Size		8G
Swap Size		2G

Coremon Start Time	2020-07-01 08:35:24 UTC
Coremon Uptime		11 days 16 hours 37 minutes
`)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
