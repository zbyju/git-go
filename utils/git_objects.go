package utils

import (
	"errors"
	"fmt"
)

func NameToPath(name string) (string, error) {
	if len(name) < 2 {
		return "", errors.New("Name of the object " + name + " cannot be converted to a path.")
	}
	prefix := name[0:2]
	suffix := name[2:]
	return fmt.Sprintf("./.git/objects/%s/%s", prefix, suffix), nil
}
