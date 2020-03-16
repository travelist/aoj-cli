package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Execute a solution with test cases",
	Run:   testCommand,
}

var testCommand = func(command *cobra.Command, args []string) {

}
