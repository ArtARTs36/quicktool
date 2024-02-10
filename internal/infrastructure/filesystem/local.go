package filesystem

import (
	"errors"
	"os"
)

type LocalFileSystem struct {
}

func (fs *LocalFileSystem) Exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}

func (fs *LocalFileSystem) GetContent(path string) ([]byte, error) {
	return os.ReadFile(path)
}
