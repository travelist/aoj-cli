package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/travelist/aoj-cli/client"
	"github.com/travelist/aoj-cli/cmd/conf"
	"net/http"
)

// default values
const defaultApiUrl = "https://judgeapi.u-aizu.ac.jp"
const defaultDataApiUrl = "https://judgedat.u-aizu.ac.jp"
const metadataFileName = "metadata.yml"

// config keys
const configKeyBaseAPIUrl = "baseUrl"
const configKeyDataAPIUrl = "dataUrl"

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
	rootCmd.PersistentFlags().StringP(configKeyBaseAPIUrl, "", defaultApiUrl, "API endpoint")
	rootCmd.PersistentFlags().StringP(configKeyDataAPIUrl, "", defaultDataApiUrl, "API endpoint")
	viper.BindPFlag(configKeyBaseAPIUrl, rootCmd.PersistentFlags().Lookup(configKeyBaseAPIUrl))
	viper.BindPFlag(configKeyDataAPIUrl, rootCmd.PersistentFlags().Lookup(configKeyDataAPIUrl))
}

// Execute executes the root command.
func Run() error {
	return rootCmd.Execute()
}

func newDefaultClient() (*client.AOJClient, error) {
	endpointURL := viper.GetString(configKeyBaseAPIUrl)
	dataEndpointURL := viper.GetString(configKeyDataAPIUrl)
	httpClient := &http.Client{}
	return client.NewClient(endpointURL, dataEndpointURL, httpClient)
}

// read and initialize configuration
func initConfig() {
	viper.AddConfigPath(conf.ConfigDirPath())
	viper.SetConfigType("toml")
	if e := viper.ReadInConfig(); e != nil {
		fmt.Printf("Warning: No configuration file. Please execute 'aoj init' to create it\n")
		return
	}
	hasReadConfigFile = true
}
