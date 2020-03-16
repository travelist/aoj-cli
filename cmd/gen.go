package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/client/response"
	"github.com/travelist/aoj-cli/common"
	"io"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

func init() {
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen <PROBLEM_ID>",
	Short: "Generate a boilerplate code and test cases",
	Run:   genCommand,
}

var genCommand = func(command *cobra.Command, args []string) {

	problemId := args[1]
	fmt.Printf("Generate files for %s...\n", problemId)
	client, e := newDefaultClient()
	if e != nil {
		fmt.Printf("Could not create API Client: %v\n", client)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	tss, e := client.FindByProblemIdSamples(ctx, args[1])
	if e != nil {
		fmt.Printf("Could not retrieve test cases: %s\n", problemId)
		os.Exit(1)
	}

	// create directory
	problemDir := filepath.Join(common.WorkspaceDirPath(), problemId)
	e = os.Mkdir(problemDir, os.ModeDir)
	if e != nil {
		fmt.Printf("Could not create a directory:  %s\n", problemDir)
		os.Exit(1)
	}

	// generate boilerplate code
	generateMetadataFile(problemDir, problemId)

	// if source code does not exist
	generateSourceCodeFile(problemDir)

	for _, ts := range tss {
		if e := generateTestCaseFiles(problemDir, ts); e != nil {
			fmt.Printf("%v\n", e)
		}
	}
}

func generateMetadataFile(problemDir string, problemId string) error {
	metadataFilePath := filepath.Join(problemDir, common.MetadataFileName)
	metadataFile, e := os.Create(metadataFilePath)
	if e != nil {
		fmt.Printf("Could not create a metadata.yml file:  %s\n", metadataFilePath)
		return e
	}
	defer metadataFile.Close()

	tmpl := template.Must(template.ParseGlob(common.MetadataFileTemplate))
	return tmpl.Execute(metadataFile, common.MetadataFileParam{ProblemId: problemId})
}

func generateTestCaseFiles(problemDir string, testCaseSample response.TestCaseSampleResponse) error {
	fmt.Printf(testCaseSample.In)
	fmt.Printf(testCaseSample.Out)
	return nil
}

func generateSourceCodeFile(problemDir string) error {
	genFileName := common.GenFileName()
	sourceFilePath := filepath.Join(problemDir, genFileName)
	sourceFile, e := os.Create(sourceFilePath)
	if e != nil {
		fmt.Printf("Could not create/open a config file at %s : %s", sourceFilePath, e.Error())
		return e
	}
	defer sourceFile.Close()

	TemplateFilePath := common.TemplateFilePath()
	if _, e := os.Stat(TemplateFilePath); os.IsNotExist(e) {
		lang := common.CodingLanguage()
		sourceFile.Write([]byte(common.LanguageToDefaultTemplate[lang]))
		fmt.Printf("Could not find a template file at %s\n", TemplateFilePath)
		return nil
	}

	templateFile, e := os.Open(TemplateFilePath)
	if e == nil {
		fmt.Printf("Could not open a template file: %s\n", TemplateFilePath)
		lang := common.CodingLanguage()
		sourceFile.Write([]byte(common.LanguageToDefaultTemplate[lang]))
		return nil
	}
	defer templateFile.Close()

	if _, e = io.Copy(sourceFile, templateFile); e != nil {
		fmt.Printf("%s\n", e.Error())
	}

	return nil
}
