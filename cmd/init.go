package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/cmd/conf"
	tmpl2 "github.com/travelist/aoj-cli/cmd/tmpl"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"text/template"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize and configure aoj-cli",
	Run:   initCommand,
}

// 1. Generate a configuration directory ("~/.aoj-cli")
// 2. Generate a configuration file ("~/.aoj-cli/config.toml")
// 3. Generate a template file ("~/.aoj-cli/template.txt")
var initCommand = func(command *cobra.Command, args []string) {
	confDir := conf.ConfigDirPath()

	if _, e := os.Stat(confDir); os.IsNotExist(e) {
		if e := os.Mkdir(confDir, 0700); e != nil {
			fmt.Printf("Could not create a config directory: %s\n", e.Error())
			return
		}
	}

	confFile := filepath.Join(confDir, conf.ConfigFileName)

	// TODO check the existence of the configuration file and ask the user whether to overwrite it or not.

	file, e := os.OpenFile(confFile, os.O_RDWR|os.O_CREATE, 0755)
	if e != nil {
		fmt.Printf("Could not create/open a config file at %s : %s\n", confFile, e.Error())
		return
	}
	defer file.Close()

	tmpl := template.Must(template.New("AOJConfig").Parse(tmpl2.ConfigFileTemplate))

	lang, e := askLanguage()
	if e != nil {
		return
	}
	username, e := askUsername()
	if e != nil {
		return
	}

	password, e := askPassword()
	if e != nil {
		return
	}

	param, _ := tmpl2.LanguageToDefaultConfigParam[lang]
	param.Username = username
	param.Password = password
	tmpl.Execute(file, param)

	templateFilePath := conf.GetGenTemplateFile()
	templateFile, e := os.OpenFile(templateFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if e != nil {
		fmt.Printf("Could not create/open a config file at %s : %s\n", templateFilePath, e.Error())
		return
	}
	defer templateFile.Close()
	templateFile.Write([]byte(tmpl2.LanguageToDefaultTemplate[lang]))

	fmt.Printf("Configuration is successfully generated under ~/.aoj-cli\n")
}

func ask(valid map[string]bool) (string, error) {
	// TODO make validator (ex. max and min length check, regex check, ... etc)
	var response string
	for true {
		if _, e := fmt.Scanln(&response); e != nil {
			fmt.Printf("Unexpected const: %s\n", e.Error())
			return "", fmt.Errorf("could not parse user input\n")
		}

		if len(valid) == 0 {
			return response, nil
		}

		if _, ok := valid[response]; ok {
			return response, nil
		}

		fmt.Printf("%s is not a valid value. Select from the following:\n", response)
		for key, _ := range valid {
			fmt.Printf("%s\n", key)
		}
	}
	return "", fmt.Errorf("unexpected const")
}

func askLanguage() (string, error) {
	validLang := map[string]bool{}

	fmt.Printf("Coding language? [")
	langs := []string{}
	for _, v := range tmpl2.ValidLanguage {
		validLang[v] = true
		langs = append(langs, v)
	}
	fmt.Printf("%s]\n", strings.Join(langs, ","))

	language, e := ask(validLang)
	if e != nil {
		return "", e
	}

	return language, nil
}

func askUsername() (string, error) {
	fmt.Printf("Username?\n")
	var response string
	if _, e := fmt.Scanln(&response); e != nil {
		fmt.Printf("Unexpected const: %s\n", e.Error())
		return "", fmt.Errorf("could not parse user input\n")
	}

	return response, nil
}

func askPassword() (string, error) {
	fmt.Printf("Password?\n")
	pswd, e := terminal.ReadPassword(syscall.Stdin)
	return string(pswd), e
}
