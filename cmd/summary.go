package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use: "summary",
	Aliases: []string{"su", "sum", "summ", "summa", "summar", "ov", "ove", "over", "overv", "overvi", "overvie", "overview"},
	Short: "Shows an overview of the system status and resources.",
	Long: `
The 'summary' command shows an overview of the system status and resources.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`UNIT		TYPE		STATUS

host01		host		OK
Uptime		host		11 days 16 hours 38 minutes
Load		host		0.77, 1.02, 1.15

Processor	system		[||||                                     10.0%]
 1		system		[|                                         2.7%]
 2		system		[|                                         1.2%]
 3		system		[||                                        4.7%]
 4		system		[|                                         1.4%]
Memory		system		[|||||||||||||||||||                 2.71G/8.0G]
Swap		system		[|                                  16.9M/2.00G]
IOWait		system		[|                                         1.4%]

eth0: in	network		[|                                      9.6Mbps]
eth0: out	network		[                                       0.5Mbps]
lo: in		network		[                                       0.0Mbps]
lo: out		network		[                                       0.0Mbps]

sda1		storage		[|||||||                                  17.2%]
sda2		storage		[||||||||||||                             26.1%]
sdb1		storage		[|||||||||||||||||||||||                  50.6%]

coremon		service		OK
openssh		service		OK
chronyd		service		OK
openntpd	service		OK
opensmtpd	service		OK
`)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
