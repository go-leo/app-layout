package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
