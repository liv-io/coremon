package cmd

import (
	"fmt"

	"github.com/liv-io/coremon/lib/coremon"
	"github.com/liv-io/coremon/lib/cpu"
	"github.com/liv-io/coremon/lib/hostname"
	"github.com/liv-io/coremon/lib/mem"
	"github.com/liv-io/coremon/lib/os"
	"github.com/liv-io/coremon/lib/sys"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"inf", "in"},
	Short:   "Shows system information",
	Long: `
The 'info' command shows information about the system.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n" + "Hostname		" + hostname.GetHostname("short"))
		fmt.Println("FQDN			" + hostname.GetHostname("long") + "\n")

		fmt.Println("IP Address		" + "10.1.1.11")
		fmt.Println("Network			" + "10.1.1.0/24")
		fmt.Println("Default Gateway		" + "10.1.1.1" + "\n")

		fmt.Println("System Start Time	" + sys.GetBoottime())
		fmt.Println("System Uptime		" + sys.GetUptime() + "\n")

		fmt.Println("OS Name			" + os.GetName())
		fmt.Println("OS Version		" + os.GetVersion())
		fmt.Println("OS Architecture		" + os.GetArchitecture() + "\n")

		fmt.Println("Processor Sockets	" + cpu.GetSockets())
		fmt.Println("Processor Cores		" + cpu.GetCores() + "\n")

		fmt.Println("Memory Size		" + mem.GetMemory())
		fmt.Println("Swap Size		" + mem.GetSwap() + "\n")

		fmt.Println("Coremon Start Time	" + coremon.GetStarttime())
		fmt.Println("Coremon Uptime		" + coremon.GetUptime() + "\n")
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
