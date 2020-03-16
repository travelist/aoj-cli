package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(submitCmd)
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a solution",
	Run:   submitCommand,
}

var submitCommand = func(command *cobra.Command, args []string) {
}
