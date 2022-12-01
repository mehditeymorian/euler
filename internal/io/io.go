package io

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var ErrDirNotExists = errors.New("directory does not exists")

func Files(dir string) ([]string, error) {
	// check if directory exist
	if !DirExists(dir) {
		return nil, ErrDirNotExists
	}

	var files []string

	walkFunc := func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}

		return nil
	}

	if err := filepath.Walk(dir, walkFunc); err != nil {
		return nil, err
	}

	return files, nil
}

func DirExists(dir string) bool {
	_, err := os.Stat(dir)

	return !os.IsNotExist(err)
}
