package conf

import "github.com/spf13/viper"

func GetSubmitSourceFileName() string {
	filename := viper.GetString("submit.source_file_name")
	if len(filename) == 0 {
		return GetGenDestinationFileName()
	}
	return filename
}
