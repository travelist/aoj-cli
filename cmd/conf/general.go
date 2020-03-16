package conf

import "github.com/spf13/viper"

func GetGeneralLanguage() string {
	return viper.GetString("general.language")
}

func GetGeneralUsername() string {
	return viper.GetString("general.username")
}

func GetGeneralPassword() string {
	return viper.GetString("general.password")
}
