package conf

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const ConfigDirName = ".aoj-cli"
const ConfigFileName = "config.toml"

var configDirPath string

func init() {
	var homeDir string
	if u, e := user.Current(); e != nil {
		homeDir = userHomeDir()
	} else {
		homeDir = u.HomeDir
	}

	configDirPath = filepath.Join(homeDir, ConfigDirName)
}

func ConfigDirPath() string {
	return configDirPath
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
