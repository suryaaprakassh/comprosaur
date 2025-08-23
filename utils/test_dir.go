package utils

import (
	"os"
	"path/filepath"
	"testing"
)

type TestDir struct {
	Path string
}

func NewTestDir(t *testing.T, dirs []string, files []string) *TestDir {
	tmp := t.TempDir()
	for _, dir := range dirs {
		path := filepath.Join(tmp, dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	for _, file := range files {
		path := filepath.Join(tmp, file)
		_, err := os.Create(path)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	return &TestDir{
		Path: tmp,
	}
}

func (t *TestDir) Get(path string) string {
	return filepath.Join(t.Path, path)
}
