package model

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"os"
	"path/filepath"

	"github.com/codecrafters-io/git-starter-go/utils"
)

type GitObject struct {
	hash [20]byte
}

func (g *GitObject) Hash() [20]byte { return g.hash }
func (g *GitObject) HashS() string  { return fmt.Sprintf("%x", g.hash) }
func (g *GitObject) Path() string {
	return utils.NameToPath(g.HashS())
}
func (g *GitObject) FullPath() string {
	return utils.PathToObjects(g.Path())
}
func (g *GitObject) compress(content string) []byte {
	var in bytes.Buffer
	buf := []byte(content)
	w := zlib.NewWriter(&in)
	w.Write(buf)
	w.Close()
	return in.Bytes()
}
func (g *GitObject) write(content string) error {
	err := os.MkdirAll(filepath.Dir(g.FullPath()), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(g.FullPath(), g.compress(content), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
