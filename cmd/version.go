package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var version string = "0.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `
The 'version' command shows detailed information about the build environment and
the version of this software.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("coremon, version: %s\n  build date: %s\n  go version: %v\n  go os     : %v\n  go arch   : %v\n",
		version, time.Now().UTC().Format(time.RFC3339), runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}
