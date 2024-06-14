package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/git-starter-go/model"
	"github.com/spf13/cobra"
)

var HashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "Save a file hashed using SHA",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf("Could not read flag: %v", err)
		}
		if path == "" {
			log.Fatalf("Path flag is required")
		}
		err = runHashObjectCommand(path)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

func init() {
	HashObjectCmd.Flags().StringP("path", "w", "", "Path to the file")
	HashObjectCmd.MarkFlagRequired("path")
}

func runHashObjectCommand(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	blobObject := model.NewBlobObject(string(content))

	err = (*blobObject).Write()
	if err != nil {
		return err
	}

	fmt.Printf("%s", blobObject.Git.HashS())

	return nil
}
