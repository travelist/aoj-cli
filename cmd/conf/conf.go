package conf

import (
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

const ConfigDirName = ".aoj-cli"
const ConfigFileName = "config.toml"
const SessionFileName = "cookie.txt"

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

func SaveSession(token string) error {
	path := filepath.Join(ConfigDirPath(), SessionFileName)
	f, e := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if e != nil {
		return e
	}
	defer f.Close()
	io.Copy(f, strings.NewReader(token))
	return nil
}

func GetSession() string {
	path := filepath.Join(ConfigDirPath(), SessionFileName)
	if _, e := os.Stat(path); os.IsNotExist(e) {
		return ""
	}

	b, e := ioutil.ReadFile(path)
	if e != nil {
		return ""
	}

	return string(b)
}
