package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/travelist/aoj-cli/client"
	"github.com/travelist/aoj-cli/common"
	"net/http"
)

const defaultApiUrl = "https://judgeapi.u-aizu.ac.jp"

var (
	// true if the configuration file has been read
	hasReadConfigFile = false

	rootCmd = &cobra.Command{
		Use:           "common",
		Short:         "A command-line tool for Aizu Online Judge (AOJ)",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("baseUrl", "", defaultApiUrl, "API endpoint")
	viper.BindPFlag("baseUrl", rootCmd.PersistentFlags().Lookup("baseUrl"))
}

// Execute executes the root command.
func Run() error {
	return rootCmd.Execute()
}

func newDefaultClient() (*client.AOJClient, error) {
	endpointURL := viper.GetString("url")
	httpClient := &http.Client{}
	return client.NewClient(endpointURL, httpClient)
}

// read and initialize configuration
func initConfig() {
	viper.AddConfigPath(common.ConfigDirPath())
	viper.SetConfigType("toml")
	if e := viper.ReadInConfig(); e != nil {
		fmt.Printf("Warning: No configuration file. Please execute 'aoj init' to create it\n")
		return
	}
	hasReadConfigFile = true
}
