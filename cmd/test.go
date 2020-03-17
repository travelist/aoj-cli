package cmd

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/cmd/conf"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Execute a solution with test cases",
	Run:   testCommand,
}

var testCommand = func(command *cobra.Command, args []string) {
	c := conf.GetTestCommand()
	if len(c) == 0 {
		fmt.Printf("Cannot find a test command. Please check %s#[test].command\n",
			filepath.Join(conf.ConfigDirPath(), conf.ConfigFileName))
		os.Exit(1)
	}

	if e := executeBeforeAll(); e != nil {
		fmt.Printf("%v\n", e)
		os.Exit(1)
	}

	testIns := listTestInputFiles()
	testOuts := listTestOutputFiles()
	if len(testIns) != len(testOuts) {
		fmt.Printf("Unmatch the number of in_{}.txt files and of out_{}.txt files\n")
		fmt.Printf("input files: %s\n", strings.Join(testIns, ","))
		fmt.Printf("output files: %s\n", strings.Join(testOuts, ","))
		os.Exit(1)
	}

	if len(testIns) == 0 {

	}

	fmt.Printf("Found %d test cases\n", len(testIns))
	for index, _ := range testIns {
		fmt.Printf("Running test set #%d\n", index)

		if e := executeBeforeEach(); e != nil {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		}

		// TODO TLE check
		output, e := executeCommand(strings.Split(c, " "), testIns[index])
		if e != nil {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		}

		success, e := checkOutput(testOuts[index], output)
		if e != nil {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		}

		if success {
			fmt.Printf("%s: test set #%d\n", color.GreenString("PASS"), index)
		} else {
			fmt.Printf("%v: test set #%d\n", color.RedString("FAIL"), index)
			fmt.Printf("[Input]\n")
			in, _ := getFileBody(testIns[index])
			fmt.Printf("%s", in)
			fmt.Printf("[Expected]\n")
			out, _ := getFileBody(testOuts[index])
			fmt.Printf("%s", color.RedString(out))
			fmt.Printf("[Actual]\n")
			fmt.Printf(color.RedString(string(output)))
		}

		if e := executeAfterEach(); e != nil {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		}
	}

	if e := executeAfterAll(); e != nil {
		fmt.Printf("%v\n", e)
		os.Exit(1)
	}

}

func executeBeforeAll() error {
	c := conf.GetTestBeforeAll()
	if len(c) == 0 {
		fmt.Printf("[%s] %s\n", color.GreenString("Before All"), color.BlueString("pass"))
		return nil
	}

	a := strings.Split(c, " ")
	fmt.Printf("[%s] %s\n", color.GreenString("Before All"), color.BlueString(c))

	var out, err bytes.Buffer
	command := exec.Command(a[0], a[1:]...)
	command.Stdout = &out
	command.Stderr = &err

	e := command.Run()

	fmt.Printf(out.String())
	fmt.Printf(color.RedString(err.String()))
	if e != nil {
		return errors.Wrap(e, "failed to execute [test].before_all command")
	}

	return nil
}

func executeBeforeEach() error {
	c := conf.GetTestBeforeEach()
	if len(c) == 0 {
		return nil
	}

	fmt.Printf("[%s] %s\n", color.GreenString("Before Each"), color.BlueString(c))
	a := strings.Split(c, " ")
	e := exec.Command(a[0], a[1:]...).Run()

	if e != nil {
		return e
	}

	return nil
}

func executeCommand(command []string, inputFilePath string) ([]byte, error) {
	fmt.Printf("[%s] %s\n", color.GreenString("Test"), color.BlueString(strings.Join(command, " ")))
	in, e := os.Open(inputFilePath)
	if e != nil {
		return []byte{}, fmt.Errorf("cannot open %s, %v", inputFilePath, e)
	}
	defer in.Close()

	c := exec.Command(command[0], command[1:]...)
	stdin, _ := c.StdinPipe()
	io.Copy(stdin, in)
	return c.Output()
}

func checkOutput(expectedOutputFilePath string, output []byte) (bool, error) {
	expected, e := os.Open(expectedOutputFilePath)
	if e != nil {
		return false, fmt.Errorf("cannot open %s: %v", expectedOutputFilePath, e)
	}
	defer expected.Close()

	exp, e := ioutil.ReadAll(expected)
	if e != nil {
		return false, fmt.Errorf("cannot read %s: %v", expectedOutputFilePath, e)
	}

	return bytes.Equal(exp, output), nil
}

func getFileBody(filePath string) (string, error) {
	in, e := os.Open(filePath)
	if e != nil {
		return "", fmt.Errorf("cannot open %s, %v", filePath, e)
	}
	defer in.Close()

	body, e := ioutil.ReadAll(in)
	if e != nil {
		return "", fmt.Errorf("cannot read %s, %v", filePath, e)
	}

	return string(body), nil
}

func executeAfterEach() error {
	c := conf.GetTestAfterEach()
	if len(c) == 0 {
		return nil
	}

	fmt.Printf("[%s] %s\n", color.GreenString("After Each"), color.BlueString(c))

	a := strings.Split(c, " ")
	e := exec.Command(a[0], a[1:]...).Run()

	if e != nil {
		return e
	}

	return nil
}

func listTestInputFiles() []string {
	path, e := os.Getwd()
	if e != nil {
		fmt.Printf("%v\n", e)
		return []string{}
	}

	var testInputFiles []string
	fileInfo, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Printf("%v\n", e)
		return []string{}
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			continue
		}

		if !strings.HasPrefix(file.Name(), "in_") {
			continue
		}

		testInputFiles = append(testInputFiles, file.Name())
	}

	return testInputFiles
}

func listTestOutputFiles() []string {
	path, e := os.Getwd()
	if e != nil {
		fmt.Printf("%v\n", e)
		return []string{}
	}

	var testInputFiles []string
	fileInfo, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Printf("%v\n", e)
		return []string{}
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			continue
		}

		if !strings.HasPrefix(file.Name(), "out_") {
			continue
		}

		testInputFiles = append(testInputFiles, file.Name())
	}

	return testInputFiles
}

func executeAfterAll() error {
	c := conf.GetTestAfterAll()
	if len(c) == 0 {
		return nil
	}
	fmt.Printf("[%s] %s\n", color.GreenString("After Each"), color.BlueString(c))

	a := strings.Split(c, " ")
	e := exec.Command(a[0], a[1:]...).Run()

	if e != nil {
		return e
	}

	return nil
}
