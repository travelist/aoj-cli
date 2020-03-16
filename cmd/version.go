package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version  string
	Revision string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run:   versionCommand,
}

var versionCommand = func(cmd *cobra.Command, args []string) {
	fmt.Printf("version: %s, revision: %s\n", Version, Revision)
}
