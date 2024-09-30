package main

import (
	"duckVault/cmd"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "duckvault",
	Version: "0.1",
	Short:   "DuckVault is a tool to convert your data into K8-Secrets",
}

func init() {
	rootCmd.AddCommand(cmd.NewCmdGenerate())
}

func main() {

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
