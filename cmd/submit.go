package cmd

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/client"
	"github.com/travelist/aoj-cli/client/request"
	"github.com/travelist/aoj-cli/client/response"
	"github.com/travelist/aoj-cli/cmd/boilerplate"
	"github.com/travelist/aoj-cli/cmd/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const defaultSubmissionURL = "https://onlinejudge.u-aizu.ac.jp/recent_judges"

func init() {
	rootCmd.AddCommand(submitCmd)
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a solution",
	Run:   submitCommand,
}

var submitCommand = func(command *cobra.Command, args []string) {

	currentDir, e := os.Getwd()
	if e != nil {
		fmt.Printf("Could not get current directory:  %v\n", e)
		os.Exit(1)
	}

	meta, e := ReadMetadata(filepath.Join(currentDir, metadataFileName))
	if e != nil {
		fmt.Printf("Could not read %s: %v\n", metadataFileName, e)
		os.Exit(1)
	}

	client, e := newDefaultClient()
	if e != nil {
		fmt.Printf("could not create a client %v\n", e)
		os.Exit(1)
	}

	if client.IsAuthenticated() {
		fmt.Printf("Loaded session from %s\n", filepath.Join(conf.ConfigDirPath(), conf.SessionFileName))
	} else {
		if e := askAuthentication(client); e != nil {
			fmt.Printf("%s : coule not authenticate user %v\n",
				color.RedString("Authentication Failed"), e)
			os.Exit(1)
		}
	}

	sourceCode, e := readSubmitFile()
	if e != nil {
		fmt.Printf("could not read source file: %v\n", e)
		os.Exit(1)
	}

	lang := conf.GetSubmitLanguage()
	if _, ok := boilerplate.ValidLanguageSet[lang]; !ok {
		fmt.Printf("%s %s %s\n",
			color.YellowString("WARNING:"),
			color.YellowString(lang),
			color.YellowString("seems invalid. Please check [submit].language configuration."))
	}

	req := request.SubmitRequest{
		ProblemID:  meta.ProblemId,
		Language:   lang,
		SourceCode: sourceCode,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, e := client.Submit(ctx, req)
	if e != nil {
		fmt.Errorf("could not submit the solution: %v\n", e)
		os.Exit(1)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	records, e := client.FindRecentSubmissionRecords(ctx)
	if e != nil {
		fmt.Printf("%s\n", color.RedString(e.Error()))
		fmt.Printf("%s Check %s\n",
			color.GreenString("DONE!"),
			color.BlueString(defaultSubmissionURL))
	} else {
		fmt.Printf("%s Check %s\n",
			color.GreenString("DONE!"),
			color.BlueString(findSubmissionURL(records, result.Token)))
	}
}

func findSubmissionURL(records response.SubmissionRecordsResponse, token string) string {
	if records == nil || len(token) == 0 {
		return defaultSubmissionURL
	}

	for _, v := range records {
		if v.Token == token {
			return fmt.Sprintf("https://onlinejudge.u-aizu.ac.jp/status/users/%s/submissions/1/%s/judge/%d/%s",
				v.UserID, v.ProblemID, v.JudgeID, v.Language)
		}
	}

	return defaultSubmissionURL
}

func askAuthentication(client *client.AOJClient) error {
	username, e := askUsername()
	if e != nil {
		return e
	}
	password, e := askPassword()
	if e != nil {
		return e
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, e = client.Login(ctx, request.LoginRequest{
		ID:       username,
		Password: password,
	})
	if e != nil {
		return e
	}

	conf.SaveSession(client.Token)
	return nil
}

func readSubmitFile() (string, error) {
	name := conf.GetSubmitSourceFileName()
	currentDir, e := os.Getwd()
	if e != nil {
		return "", fmt.Errorf("Could not get current directory:  %v\n", e)
	}
	path := filepath.Join(currentDir, name)
	content, e := ioutil.ReadFile(path)
	return string(content), e
}

type ProblemMetadata struct {
	ProblemId string `yaml:"problem_id"`
}

func ReadMetadata(path string) (*ProblemMetadata, error) {
	yamlFile, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, fmt.Errorf("could not read %s: %v ", path, e)
	}

	var metadata ProblemMetadata
	e = yaml.Unmarshal(yamlFile, &metadata)
	if e != nil {
		return nil, fmt.Errorf("could not parse %s: %v ", path, e)
	}

	return &metadata, nil
}
