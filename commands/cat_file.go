package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/git-starter-go/model"
	"github.com/codecrafters-io/git-starter-go/utils"
	"github.com/spf13/cobra"
)

var CatFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Print the contents of a file",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf("Could not read flag: %v", err)
		}
		if path == "" {
			log.Fatalf("Path flag is required")
		}
		err = runCatFileCommand(path)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

func init() {
	CatFileCmd.Flags().StringP("path", "p", "", "Name of the git object")
	CatFileCmd.MarkFlagRequired("path")
}

func runCatFileCommand(name string) error {
	path := utils.PathToObjects(utils.NameToPath(name))

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	blobObject, err := model.NewBlobObjectFromFile(file)
	if err != nil {
		return err
	}

	fmt.Printf("%s", blobObject.Content())

	return nil
}
