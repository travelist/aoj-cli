package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/common"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"path/filepath"
	"syscall"
	"text/template"
)

func init() {
	rootCmd.AddCommand(InitCmd)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize and configure common-cli",

	// 1. Generate a configuration directory ("~/.aoj-cli")
	// 2. Generate a configuration file ("~/.aoj-cli/config.toml")
	// 3. Generate a template file ("~/.aoj-cli/template.txt")
	Run: func(command *cobra.Command, args []string) {
		confDir := common.ConfigDirPath()

		if e := os.Mkdir(confDir, os.ModeDir); e != nil {
			fmt.Printf("Could not create a config directory: %s", e.Error())
			return
		}

		confFile := filepath.Join(confDir, "config.toml")

		// TODO check the existence of the configuration file and ask the user whether to overwrite it or not.

		file, e := os.OpenFile(confFile, os.O_RDWR|os.O_CREATE, 0755)
		if e != nil {
			fmt.Printf("Could not create/open a config file at %s : %s", confFile, e.Error())
			return
		}
		defer file.Close()

		tmpl := template.Must(template.ParseGlob(common.ConfigFileTemplate))

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

		param, _ := common.LanguageToDefaultConfigParam[lang]
		param.Username = username
		param.Password = password
		tmpl.Execute(file, param)

		templateFilePath := common.TemplateFilePath()
		templateFile, e := os.OpenFile(templateFilePath, os.O_RDWR|os.O_CREATE, 0755)
		if e != nil {
			fmt.Printf("Could not create/open a config file at %s : %s", templateFilePath, e.Error())
			return
		}
		defer templateFile.Close()
		templateFile.Write([]byte(common.LanguageToDefaultTemplate[lang]))
	},
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

	fmt.Printf("What is main language? (options)\n")
	for _, v := range common.ValidLanguage {
		validLang[v] = true
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("\n")

	language, e := ask(validLang)
	if e != nil {
		return "", e
	}

	return language, nil
}

func askUsername() (string, error) {
	fmt.Printf("What is username?\n")
	var response string
	if _, e := fmt.Scanln(&response); e != nil {
		fmt.Printf("Unexpected const: %s\n", e.Error())
		return "", fmt.Errorf("could not parse user input\n")
	}

	return response, nil
}

func askPassword() (string, error) {
	fmt.Printf("What is password?\n")
	pswd, e := terminal.ReadPassword(syscall.Stdin)
	return string(pswd), e
}
