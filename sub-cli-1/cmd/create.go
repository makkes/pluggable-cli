package cmd

import "github.com/spf13/cobra"

func CreateCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use: "create",
	}

	createCmd.AddCommand(CreatePipelineCmd())

	return createCmd
}
