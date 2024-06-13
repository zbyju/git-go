package commands

import (
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/codecrafters-io/git-starter-go/model"
	"github.com/codecrafters-io/git-starter-go/utils"
	"github.com/spf13/cobra"
)

var LsTreeCmd = &cobra.Command{
	Use:   "ls-tree",
	Short: "Prints a tree object",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		nameOnly, err := cmd.Flags().GetBool("name-only")
		if err != nil {
			log.Fatalf("Could not read flag: %v", err)
		}

		err = runLsTreeCommand(filePath, nameOnly)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

func init() {
	LsTreeCmd.Flags().Bool("name-only", true, "Prints only the names of the directories and files.")
}

func runLsTreeCommand(name string, nameOnly bool) error {
	path, err := utils.NameToPath(name)
	if err != nil {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader, err := zlib.NewReader(file)
	if err != nil {
		return err
	}
	defer reader.Close()

	p, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	tree, err := model.ParseTreeObject(p)

	if err != nil {
		return err
	}

	fmt.Printf("%s", tree.ToString(nameOnly))

	return nil
}
