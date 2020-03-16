package conf

import (
	"github.com/spf13/viper"
	"path/filepath"
)

const defaultTemplateFileName = "template.txt"

func GetGenTemplateFile() string {
	a := viper.GetString("gen.template_file")
	if len(a) == 0 {
		return defaultTemplateFile()
	}
	return a
}

func GetGenDestinationFileName() string {
	return viper.GetString("gen.destination_file_name")
}

func defaultTemplateFile() string {
	return filepath.Join(ConfigDirPath(), defaultTemplateFileName)
}
