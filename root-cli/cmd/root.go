package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.e13.dev/sub-cli-1/cmd"
)

func Execute() {
	rootCmd := &cobra.Command{}

	subCLI1RootCmd := cmd.RootCommand()
	for _, subCmd := range subCLI1RootCmd.Commands() {
		rootCmd.AddCommand(subCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
