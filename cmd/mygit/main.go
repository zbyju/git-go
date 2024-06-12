package main

import (
	"os"

	"github.com/codecrafters-io/git-starter-go/commands"
	"github.com/spf13/cobra"
)

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	rootCmd := &cobra.Command{Use: "myapp"}
	rootCmd.AddCommand(commands.InitCmd)
	rootCmd.AddCommand(commands.CatFileCmd)
	rootCmd.AddCommand(commands.HashObjectCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
