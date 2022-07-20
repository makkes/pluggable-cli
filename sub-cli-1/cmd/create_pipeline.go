package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CreatePipelineCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pipeline",
		Short: "Creates a new deployment pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("created new pipeline")
		},
	}
}
