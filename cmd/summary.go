package cmd

import (
	"fmt"

	"github.com/liv-io/coremon/lib/hostname"
        "github.com/liv-io/coremon/lib/sys"

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
		fmt.Println("\n" + hostname.GetHostname("short")+ "		" + "host" + "		" + "OK")
		fmt.Println("Uptime" + "		" + "host" + "		" + sys.GetUptime())
		fmt.Println("Load" + "		" + "host" + "		" + "0.77, 1.02, 1.15" + "\n")

		fmt.Println("Processor" + "	" + "system" + "		" + "[||||                                     10.0%]")
		fmt.Println("  1" + "		" + "system" + "		" + "[|                                         2.7%]")
		fmt.Println("  2" + "		" + "system" + "		" + "[|                                         1.2%]")
		fmt.Println("  3" + "		" + "system" + "		" + "[||                                        4.7%]")
		fmt.Println("  4" + "		" + "system" + "		" + "[|                                         1.4%]")
		fmt.Println("Memory" + "		" + "system" + "		" + "[|||||||||||||||||||                 2.71G/8.0G]")
		fmt.Println("Swap" + "		" + "system" + "		" + "[|                                  16.9M/2.00G]")
		fmt.Println("IOWait" + "		" + "system" + "		" + "[|                                         1.4%]" + "\n")

		fmt.Println("eth0: in" + "	" + "network" + "		" + "[|                                      9.6Mbps]")
		fmt.Println("eth0: out" + "	" + "network" + "		" + "[|                                      0.5Mbps]")
		fmt.Println("lo: in" + "		" + "network" + "		" + "[|                                      0.0Mbps]")
		fmt.Println("lo: out" + "		" + "network" + "		" + "[|                                      0.0Mbps]" + "\n")

		fmt.Println("sda1" + "		" + "storage" + "		" + "[|||||||                                  17.2%]")
		fmt.Println("sda2" + "		" + "storage" + "		" + "[||||||||||||                             26.1%]")
		fmt.Println("sdb1" + "		" + "storage" + "		" + "[|||||||||||||||||||||||                  50.6%]" + "\n")

		fmt.Println("coremon" + "		" + "service" + "		" + "OK")
		fmt.Println("openssh" + "		" + "service" + "		" + "OK")
		fmt.Println("chronyd" + "		" + "service" + "		" + "OK")
		fmt.Println("openntpd" + "	" + "service" + "		" + "OK")
		fmt.Println("opensmtpd" + "	" + "service" + "		" + "OK" + "\n")
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
