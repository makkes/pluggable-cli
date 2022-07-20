package main

import (
	"os"

	"go.e13.dev/sub-cli-1/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		os.Exit(1)
	}

}
