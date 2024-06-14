package utils

import (
	"os"
)

type FileInfo struct {
	Name    string
	Mode    string
	Content string
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func ListDirectories(dirPath string) ([]string, error) {
	var directories []string

	// Open the directory
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read the directory entries
	entries, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() == ".git" {
				continue
			}
			directories = append(directories, entry.Name())
		}
	}

	return directories, nil
}

func ListFiles(dirPath string) ([]FileInfo, error) {
	var files []FileInfo

	// Open the directory
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read the directory entries
	entries, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			filePath := dirPath + "/" + entry.Name()
			content, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}

			files = append(files, FileInfo{
				Name:    entry.Name(),
				Mode:    "100644",
				Content: string(content),
			})
		}
	}

	return files, nil
}
