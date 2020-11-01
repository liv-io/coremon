package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "coremon",
	Short: "Coremon monitors files, processes, sockets, network interfaces and other devices.",
	Long: `
  _________  ________  ________  ____  ____
 / ___/ __ \/ ___/ _ \/ __  __ \/ __ \/ __ \
/ /__/ /_/ / /  /  __/ / / / / / /_/ / / / /
\___/\____/_/   \___/_/ /_/ /_/\____/_/ /_/

Coremon is an open-source monitoring tool designed for the Unix console.
It aims to provide an overview of the system status and resources. Coremon
monitors files, processes, sockets, network interfaces and other devices. The
software is written in Go and licensed under the Simplified BSD License.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coremon.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".coremon" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".coremon")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
