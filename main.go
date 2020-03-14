package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/travelist/aoj-cli/cmd"
)

func mani() {
	viper.AddConfigPath("$HOME/.aoj/config")
	viper.SetConfigType("toml")

	if e := viper.ReadInConfig(); e != nil {
		if _, ok := e.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Creating a config file at: $HOME/.aoj/config.toml")
		}
	}

	cmd.Run()
}
