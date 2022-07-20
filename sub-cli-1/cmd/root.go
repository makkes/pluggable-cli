package cmd

import (
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(CreateCmd())
	return rootCmd
}
