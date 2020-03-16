package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/travelist/aoj-cli/common"
	"github.com/travelist/aoj-cli/cmd"
	"os"
)

func mani() {
	viper.AddConfigPath(common.ConfigDirPath())
	viper.SetConfigType("toml")

	if e := viper.ReadInConfig(); e != nil {
		if _, ok := e.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Could not find common-cli config. Maybe, you need to execute:\n\n")
			fmt.Printf("	common init\n\n")
			os.Exit(1)
		}
	}

	cmd.Run()
}
