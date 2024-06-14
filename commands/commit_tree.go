package commands

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/codecrafters-io/git-starter-go/model"
	"github.com/spf13/cobra"
)

var CommitTreeCmd = &cobra.Command{
	Use:   "commit-tree [tree hash]",
	Short: "Create a new commit object",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		treeHash := args[0]
		parents, err := cmd.Flags().GetStringArray("parents")
		if err != nil {
			log.Fatalf("Could not read flag: %v", err)
			os.Exit(1)
		}
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			log.Fatalf("Could not read flag: %v", err)
			os.Exit(1)
		}
		err = runCommitTreeCommand(treeHash, parents, message)
		if err != nil {
			log.Fatalf("There has been a problem running commit-tree: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	CommitTreeCmd.Flags().StringArrayP("parents", "p", []string{}, "Parents SHA")
	CommitTreeCmd.Flags().StringP("message", "m", "", "Commit message")
}

func runCommitTreeCommand(treeHash string, parents []string, message string) error {
	author := model.GitAuthor{Name: model.DEFAULT_NAME, Email: model.DEFAULT_EMAIL, Timestamp: time.Now()}
	commiter := model.GitAuthor{Name: model.DEFAULT_NAME, Email: model.DEFAULT_EMAIL, Timestamp: time.Now()}

	commitObject := model.NewCommitObject(author, commiter, treeHash, parents, message)

	commitObject.Write()

	fmt.Printf("%s", commitObject.Git.HashS())

	return nil
}
