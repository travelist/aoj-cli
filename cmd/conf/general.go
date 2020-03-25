package conf

import "github.com/spf13/viper"

func GetGeneralLanguage() string {
	return viper.GetString("general.language")
}
