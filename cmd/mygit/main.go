package main

import (
	"os"

	"github.com/codecrafters-io/git-starter-go/commands"
	"github.com/spf13/cobra"
)

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	rootCmd := &cobra.Command{Use: "git"}
	rootCmd.AddCommand(commands.InitCmd)
	rootCmd.AddCommand(commands.CatFileCmd)
	rootCmd.AddCommand(commands.HashObjectCmd)
	rootCmd.AddCommand(commands.LsTreeCmd)
	rootCmd.AddCommand(commands.WriteTreeCmd)
	rootCmd.AddCommand(commands.CommitTreeCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
