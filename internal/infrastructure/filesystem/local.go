package filesystem

import (
	"errors"
	"os"
)

type LocalFileSystem struct {
}

func (fs *LocalFileSystem) Exists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

func (fs *LocalFileSystem) GetContent(path string) ([]byte, error) {
	return os.ReadFile(path)
}
