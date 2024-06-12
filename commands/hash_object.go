package commands

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codecrafters-io/git-starter-go/utils"
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

	toHash := fmt.Sprintf("blob %d\x00%s", len(content), content)
	sha := (sha1.Sum([]byte(toHash)))
	hashed := fmt.Sprintf("%x", sha)

	hashPath, err := utils.NameToPath(hashed)
	if err != nil {
		return err
	}

	var in bytes.Buffer
	b := []byte(toHash)
	w := zlib.NewWriter(&in)
	w.Write(b)
	w.Close()

	err = os.MkdirAll(filepath.Dir(hashPath), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(hashPath, in.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Printf("%s", hashed)

	return nil
}
