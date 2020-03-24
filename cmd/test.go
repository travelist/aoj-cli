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
		fmt.Printf("%s - No test cases found\n", color.GreenString("AC"))
		os.Exit(0)
	}

	allSuccess := true

	for index, _ := range testIns {
		if e := executeBeforeEach(); e != nil {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		}

		// TODO TLE check
		out, e := executeCommand(strings.Split(c, " "), testIns[index])
		if e != nil {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		}

		success, e := checkOutput(testOuts[index], out)
		if e != nil {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		}

		if success {
			fmt.Printf("# in_%d.txt ... %s\n", index+1, color.GreenString("AC"))
		} else {
			allSuccess = false
			fmt.Printf("# in_%d.txt ... %s\n", index+1, color.RedString("WA"))

			fmt.Printf("%s\n", color.MagentaString("[Input]"))
			in, _ := getFileBody(testIns[index])
			fmt.Printf(in)

			fmt.Printf("%s\n", color.MagentaString("[Expected]"))
			expected, _ := getFileBody(testOuts[index])
			fmt.Printf(expected)

			fmt.Printf("%s\n", color.MagentaString("[Actual]"))
			fmt.Printf(string(out))
			fmt.Println()
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

	if allSuccess {
		fmt.Printf("%s\n", color.GreenString("All AC!"))
	} else {
		fmt.Printf("%s\n", color.RedString("Some cases are WRONG ANSWER"))
	}

}

func executeBeforeAll() error {
	c := conf.GetTestBeforeAll()
	if len(c) == 0 {
		return nil
	}

	a := strings.Split(c, " ")

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
