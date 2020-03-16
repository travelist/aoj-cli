package conf

import "github.com/spf13/viper"

func GetTestBeforeAll() string {
	return viper.GetString("test.before_all")
}

func GetTestBeforeEach() string {
	return viper.GetString("test.before_each")
}

func GetTestCommand() string {
	return viper.GetString("test.command")
}

func GetTestAfterEach() string {
	return viper.GetString("test.after_each")
}

func GetTestAfterAll() string {
	return viper.GetString("test.after_all")
}
