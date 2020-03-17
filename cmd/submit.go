package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/client/request"
	"github.com/travelist/aoj-cli/cmd/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

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
		fmt.Errorf("could not create a client %v\n", e)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, e = client.Login(ctx, request.LoginRequest{
		ID:       conf.GetGeneralUsername(),
		Password: conf.GetGeneralPassword(),
	})

	if e != nil {
		fmt.Errorf("authentication failed: %v\n", e)
		os.Exit(1)
	}

	sourceCode, e := readSubmitFile()
	if e != nil {
		fmt.Printf("could not read source file: %v\n", e)
		os.Exit(1)
	}

	req := request.SubmitRequest{
		ProblemID:  meta.ProblemId,
		Language:   conf.GetGeneralLanguage(),
		SourceCode: sourceCode,
	}

	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, e := client.Submit(ctx, req)
	if e != nil {
		fmt.Errorf("could not submit the solution: %v\n", e)
		os.Exit(1)
	}

	fmt.Printf("Submission Success: %s\n", result.Token)
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
	e = yaml.Unmarshal(yamlFile, metadata)
	if e != nil {
		return nil, fmt.Errorf("could not parse %s: %v ", path, e)
	}

	return &metadata, nil
}
