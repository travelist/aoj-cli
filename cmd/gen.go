package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/travelist/aoj-cli/aoj"
	"os"
)

func init() {
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a boilerplate code and test cases",
	Run: func(command *cobra.Command, args []string) {

		problemId := args[1]
		fmt.Printf("Generate files for %s...\n", problemId)
		client := aoj.NewAPIClient()
		ts, e := client.FetchTestCases(args[1])
		if e != nil {
			fmt.Printf("Could not retrieve test cases: %s\n", problemId)
			os.Exit(1)
		}

		// create directory
		e = os.Mkdir("")
		if e != nil {
			fmt.Printf("Could not create a directory:  %s\n", problemId)
			os.Exit(1)
		}

		// generate boilerplate code

		// generate test cases

	},
}
