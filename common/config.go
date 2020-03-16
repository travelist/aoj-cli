package common

import (
	"github.com/spf13/viper"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const ConfigDirName = ".aoj-cli"
const ConfigFileName = "config.toml"
const TemplateFileName = "template.txt"
const MetadataFileName = "metadata.yml"

func WorkspaceDirPath() string {
	return viper.GetString("gen.workspace_directory")
}

func CodingLanguage() string {
	return viper.GetString("general.language")
}

func SubmitFileName() string {
	return viper.GetString("submit.source_file_name")
}

func GenFileName() string {
	return viper.GetString("gen.source_file_name")
}

func TemplateFilePath() string {
	return filepath.Join(ConfigDirPath(), TemplateFileName)
}

var configDirPath string

func ConfigDirPath() string {
	return configDirPath
}

func init() {
	var homeDir string
	if user, e := user.Current(); e != nil {
		homeDir = userHomeDir()
	} else {
		homeDir = user.HomeDir
	}

	configDirPath = filepath.Join(homeDir, ConfigDirName)
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
