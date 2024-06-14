package model

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"sort"
	"strconv"
)

type TreeObject struct {
	entries []TreeObjectEntry
	Git     GitObject
}

type TreeObjectEntry struct {
	Mode uint32
	Name string
	Hash [20]byte
}

func (t *TreeObject) Entries() []TreeObjectEntry {
	return t.entries
}

func (t *TreeObject) AddEntry(entry TreeObjectEntry) {
	index := sort.Search(len(t.entries), func(i int) bool {
		if t.entries[i].Name == entry.Name {
			return bytes.Compare(t.entries[i].Hash[:], entry.Hash[:]) >= 0
		}
		return t.entries[i].Name >= entry.Name
	})

	t.entries = append(t.entries, TreeObjectEntry{})
	copy(t.entries[index+1:], t.entries[index:])
	t.entries[index] = entry

	t.Git.hash = sha1.Sum([]byte(t.FullContent()))
}

func (t *TreeObject) ToString(namesOnly bool) string {
	res := ""
	if namesOnly {
		for _, e := range t.entries {
			res += e.Name + "\n"
		}
	} else {
		for _, e := range t.entries {
			res += fmt.Sprintf("%d %s %s\n", e.Mode, e.Name, e.Hash)
		}
	}
	return res
}

func (t *TreeObject) Content() string {
	content := ""
	for _, entry := range t.entries {
		content += fmt.Sprintf("%d %s\x00%s", entry.Mode, entry.Name, entry.Hash)
	}
	return content
}

func (t *TreeObject) FullContent() string {
	content := t.Content()
	header := fmt.Sprintf("tree %d\x00", len(content))
	return header + content
}

func (t *TreeObject) Compress() []byte {
	return t.Git.compress(t.FullContent())
}

func (t *TreeObject) Write() error {
	return t.Git.write(t.FullContent())
}

func ParseTreeObject(data []byte) (*TreeObject, error) {
	// Find the first null character to extract the size part
	nullIdx := bytes.IndexByte(data, 0)
	if nullIdx == -1 {
		return nil, fmt.Errorf("invalid input format: no null terminator found for size")
	}

	// Initialize the TreeObject
	tree := &TreeObject{}

	// Process the remaining data to extract entries
	remainingData := data[nullIdx+1:]

	for len(remainingData) > 0 {
		// Find the next null character to separate mode/name from hash
		nullIdx = bytes.IndexByte(remainingData, 0)
		if nullIdx == -1 || len(remainingData)-nullIdx-1 < 20 {
			return nil, fmt.Errorf("invalid entry format")
		}

		// Extract the mode and name part
		modeNamePart := remainingData[:nullIdx]
		remainingData = remainingData[nullIdx+1:]

		// Find the last space to separate mode and name
		lastSpaceIdx := bytes.LastIndexByte(modeNamePart, ' ')
		if lastSpaceIdx == -1 {
			return nil, fmt.Errorf("invalid mode/name format")
		}

		modeStr := string(modeNamePart[:lastSpaceIdx])
		name := string(modeNamePart[lastSpaceIdx+1:])

		// Parse the mode
		mode64, err := strconv.ParseUint(modeStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid mode: %v", err)
		}

		// Extract the hash
		if len(remainingData) < 20 {
			return nil, fmt.Errorf("invalid hash length")
		}
		var hash [20]byte
		copy(hash[:], remainingData[:20])
		remainingData = remainingData[20:]

		// Create the TreeObjectEntry
		entry := TreeObjectEntry{
			Mode: uint32(mode64),
			Name: name,
			Hash: hash,
		}

		// Add the entry to the tree using AddEntry method
		tree.AddEntry(entry)
	}

	return tree, nil
}
