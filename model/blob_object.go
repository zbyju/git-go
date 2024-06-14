package model

import (
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

type BlobObject struct {
	Git     GitObject
	content string
}

func NewBlobObject(data string) *BlobObject {
	content := data
	if headerIndex := strings.IndexByte(data, 0); len(data) >= 5 && data[0:4] == "blob" && headerIndex != -1 {
		content = data[headerIndex+1:]
	}

	toHash := fmt.Sprintf("blob %d\x00%s", len(content), content)
	sha := sha1.Sum([]byte(toHash))

	return &BlobObject{content: content, Git: GitObject{hash: sha}}
}

func NewBlobObjectFromFile(file *os.File) (*BlobObject, error) {
	reader, err := zlib.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	p, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	str := string(p)

	return NewBlobObject(str), nil
}

func (b *BlobObject) Content() string {
	return b.content
}

func (b *BlobObject) Length() int {
	return len(b.content)
}

func (b *BlobObject) FullContent() string {
	return fmt.Sprintf("blob %d\x00%s", b.Length(), b.Content())
}

func (b *BlobObject) Compress() []byte {
	return b.Git.compress(b.FullContent())
}

func (b *BlobObject) Write() error {
	return b.Git.write(b.FullContent())
}
