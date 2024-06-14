package utils

import (
	"path/filepath"
)

func NameToPath(name string) string {
	prefix := name[0:2]
	suffix := name[2:]
	return "/" + prefix + "/" + suffix
}

func PathToObjects(path string) string {
	return filepath.Join("./.git/objects/", path)
}
