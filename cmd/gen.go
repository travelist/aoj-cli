package cmd

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/client/response"
	tmpl2 "github.com/travelist/aoj-cli/cmd/boilerplate"
	"github.com/travelist/aoj-cli/cmd/conf"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	if len(args) < 1 {
		fmt.Printf("Usage: aoj gen <PROBLEM_ID>\n")
		os.Exit(1)
	}

	problemId := args[0]
	client, e := newDefaultClient()

	if e != nil {
		fmt.Printf("Could not create API Client: %v\n", client)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tss, e := client.FindByProblemIdSamples(ctx, problemId)
	if e != nil {
		fmt.Printf("Could not retrieve test cases: %v\n", e)
		os.Exit(1)
	}

	currentDir, e := os.Getwd()
	if e != nil {
		fmt.Printf("Could not get current directory:  %v\n", e)
		os.Exit(1)
	}

	problemDir := filepath.Join(currentDir, problemId)
	if _, e := os.Stat(problemDir); os.IsNotExist(e) {
		e = os.Mkdir(problemDir, 0700)
		if e != nil {
			fmt.Printf("Could not create a directory:  %s\n", problemDir)
			os.Exit(1)
		}
	}

	for _, ts := range tss {
		if e := generateTestCaseFiles(problemDir, ts); e != nil {
			fmt.Printf("%v\n", e)
		}
	}

	// generate metadata file
	if e := generateMetadataFile(problemDir, problemId); e != nil {
		fmt.Printf("%v\n", e)
		os.Exit(1)
	}

	genFileName := conf.GetGenDestinationFileName()
	sourceFilePath := filepath.Join(problemDir, genFileName)
	fmt.Println(genFileName)
	fmt.Println(sourceFilePath)

	_, e = os.Stat(sourceFilePath)
	fmt.Println(e)
	if os.IsNotExist(e) {
		if e := generateSourceCodeFile(sourceFilePath); e != nil {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		}
	}

	fmt.Printf("%s %s %s\n",
		color.GreenString("Successfully created"),
		color.GreenString(problemId),
		color.GreenString("directory"),
	)
}

func generateMetadataFile(problemDir string, problemId string) error {
	metadataFilePath := filepath.Join(problemDir, metadataFileName)
	metadataFile, e := os.Create(metadataFilePath)
	if e != nil {
		fmt.Printf("Could not create a metadata.yml file:  %s\n", metadataFilePath)
		return e
	}
	defer metadataFile.Close()
	tp := template.Must(template.New("metadata").Parse(tmpl2.MetadataFileTemplate))
	return tp.Execute(metadataFile, tmpl2.MetadataFileParam{ProblemId: problemId})
}

func generateTestCaseFiles(problemDir string, testCaseSample response.TestCaseSampleResponse) error {
	a := fmt.Sprintf("in_%d.txt", testCaseSample.Serial)
	b := fmt.Sprintf("out_%d.txt", testCaseSample.Serial)
	inFile := filepath.Join(problemDir, a)
	outFile := filepath.Join(problemDir, b)
	inf, e := os.OpenFile(inFile, os.O_RDWR|os.O_CREATE, 0600)
	if e != nil {
		return e
	}
	defer inf.Close()
	io.Copy(inf, strings.NewReader(testCaseSample.In))

	outf, e := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, 0600)
	if e != nil {
		return e
	}
	defer outf.Close()
	io.Copy(outf, strings.NewReader(testCaseSample.Out))

	return nil
}

func generateSourceCodeFile(sourceFilePath string) error {
	sourceFile, e := os.OpenFile(sourceFilePath, os.O_RDWR|os.O_CREATE, 0600)
	if e != nil {
		fmt.Printf("Could not create/open a config file at %s : %s\n", sourceFilePath, e.Error())
		return e
	}
	defer sourceFile.Close()

	TemplateFilePath := conf.GetGenTemplateFile()
	if _, e := os.Stat(TemplateFilePath); os.IsNotExist(e) {
		lang := conf.GetGeneralLanguage()
		sourceFile.Write([]byte(tmpl2.LanguageToDefaultTemplate[lang]))
		fmt.Printf("Could not find a template file at %s\n", TemplateFilePath)
		return nil
	}

	templateFile, e := os.OpenFile(TemplateFilePath, os.O_RDWR|os.O_CREATE, 0600)
	if e == nil {
		fmt.Printf("Could not open a template file: %s\n", TemplateFilePath)
		lang := conf.GetGeneralLanguage()
		sourceFile.Write([]byte(tmpl2.LanguageToDefaultTemplate[lang]))
		return nil
	}
	defer templateFile.Close()

	if _, e = io.Copy(sourceFile, templateFile); e != nil {
		fmt.Printf("%s\n", e.Error())
	}

	return nil
}
