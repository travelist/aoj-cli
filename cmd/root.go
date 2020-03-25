package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/travelist/aoj-cli/client"
	"github.com/travelist/aoj-cli/cmd/conf"
	"net/http"
	"path/filepath"
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
		Use:           "aoj",
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
	return client.NewClient(endpointURL, dataEndpointURL, httpClient, conf.GetSession())
}

// read and initialize configuration
func initConfig() {
	confPath := filepath.Join(conf.ConfigDirPath(), "config.toml")
	viper.SetConfigFile(confPath)
	viper.SetConfigType("toml")
	if e := viper.ReadInConfig(); e != nil {
		fmt.Printf("[%s] Invalid configuration. Please check %s or execute '%s' to initialise it: %v\n",
			color.YellowString("WARNING"),
			color.BlueString("confPath"),
			color.RedString("aoj init"),
			e,
		)
		return
	}

	hasReadConfigFile = true
}
