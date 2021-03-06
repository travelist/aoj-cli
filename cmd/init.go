package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/cmd/boilerplate"
	"github.com/travelist/aoj-cli/cmd/conf"
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
	Short: "Setup aoj-cli",
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

	confFilePath := filepath.Join(confDir, conf.ConfigFileName)

	// TODO check the existence of the configuration file and ask the user whether to overwrite it or not.

	file, e := os.OpenFile(confFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if e != nil {
		fmt.Printf("Could not create/open a config file at %s : %s\n", confFilePath, e.Error())
		return
	}
	defer file.Close()

	tp := template.Must(template.New("AOJConfig").Parse(boilerplate.ConfigFileTemplate))

	lang, e := askLanguage()
	if e != nil {
		return
	}

	param, _ := boilerplate.LanguageToDefaultConfigParam[lang]
	tp.Execute(file, param)

	templateFilePath := conf.GetGenTemplateFile()
	templateFile, e := os.OpenFile(templateFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if e != nil {
		fmt.Printf("Could not create/open a config file at %s : %s\n", templateFilePath, e.Error())
		return
	}
	defer templateFile.Close()
	templateFile.Write([]byte(boilerplate.LanguageToDefaultTemplate[lang]))

	fmt.Printf("AOJ CLI is successfully initialised. Check %s\n", color.GreenString(confDir))
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

		fmt.Printf("%s is not a valid value. Select from the following:\n", color.RedString(response))
		langs := []string{}
		for key, _ := range valid {
			langs = append(langs, color.GreenString(key))
		}
		fmt.Printf("[%s]\n", strings.Join(langs, ","))

	}
	return "", fmt.Errorf("unexpected const")
}

func askLanguage() (string, error) {

	fmt.Printf("Coding language? [")
	langs := []string{}
	for _, v := range boilerplate.ValidLanguageList {
		langs = append(langs, color.GreenString(v))
	}
	fmt.Printf("%s]\n", strings.Join(langs, ","))

	language, e := ask(boilerplate.ValidLanguageSet)
	if e != nil {
		return "", e
	}

	return language, nil
}

func askUsername() (string, error) {
	fmt.Printf("Username: ")
	var response string
	if _, e := fmt.Scanln(&response); e != nil {
		fmt.Printf("Unexpected const: %s\n", e.Error())
		return "", fmt.Errorf("could not parse user input\n")
	}

	return response, nil
}

func askPassword() (string, error) {
	fmt.Printf("Password: ")
	pswd, e := terminal.ReadPassword(int(syscall.Stdin)) // Windows require casting stdin into int type.
	return string(pswd), e
}
