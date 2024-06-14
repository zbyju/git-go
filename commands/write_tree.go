package commands

import (
	"fmt"
	"path/filepath"

	"github.com/codecrafters-io/git-starter-go/model"
	"github.com/codecrafters-io/git-starter-go/utils"
	"github.com/spf13/cobra"
)

var WriteTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "Create tree objects from all files",
	Run: func(cmd *cobra.Command, args []string) {
		err := runWriteTreeCommand("./")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

func runWriteTreeCommand(path string) error {
	tree, err := createTreeObject(path)
	if err != nil {
		return err
	}

	fmt.Printf("%s", tree.Git.HashS())
	return nil
}

func createTreeObject(path string) (*model.TreeObject, error) {
	treeObject := model.TreeObject{}

	files, err := utils.ListFiles(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		pathFile := filepath.Join(path, file.Name)
		exists := utils.FileExists(pathFile)

		if !exists {
			runHashObjectCommand(pathFile)
		}

		blob := model.NewBlobObject(file.Content)
		entry := model.TreeObjectEntry{Mode: file.Mode, Name: file.Name, Hash: blob.Git.Hash()}
		treeObject.AddEntry(entry)
	}

	dirs, err := utils.ListDirectories(path)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		tree, err := createTreeObject(filepath.Join(path, dir.Name))
		if err != nil {
			return nil, err
		}
		entry := model.TreeObjectEntry{Mode: dir.Mode, Name: dir.Name, Hash: tree.Git.Hash()}

		treeObject.AddEntry(entry)
	}

	err = treeObject.Write()
	if err != nil {
		return nil, err
	}

	return &treeObject, nil
}
