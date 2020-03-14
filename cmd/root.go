package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "aoj",
		Short: "A command-line tool for Aizu Online Judge (AOJ)",
	}
)

// Execute executes the root command.
func Run() error {
	return rootCmd.Execute()
}
