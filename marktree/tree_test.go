package marktree

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	tree := NewTree()
	paths := []string{
		"/foo/boo",
		"/foo/bar/baz",
	}

	for _,path := range paths {
		tree.ToggleDir(path)
	}
	for _,path := range paths {
		assert.Equal(t,tree.IsMarked(path),true)
	}
	tree.ToggleDir("/foo/bar/a.txt")
	for _,path := range paths {
		tree.ToggleDir(path)
	}
	for _,path := range paths {
		assert.Equal(t,tree.IsMarked(path),false)
	}
	assert.Equal(t,tree.IsMarked("/foo/bar/a.txt"),true)
}


func TestMarkDrop(t *testing.T) {
	tree := NewTree()
	paths := []string{
		"/foo/bar/baz",
		"/foo/bar",
	}

	for _,path := range paths {
		tree.ToggleDir(path)
	}
	for _,path := range paths {
		assert.Equal(t,tree.IsMarked(path),true)
	}
	tree.ToggleDir("/foo")
	assert.Equal(t,tree.IsMarked("/foo/bar/a.txt"),true)
}

func TestInternalDrop(t *testing.T) {
	tree := NewTree()
	tmpDir := t.TempDir()

	fileNames := []string{"a","b","c"}

	for _,name:= range fileNames{
		fullPath := filepath.Join(tmpDir,name)	
		err := os.WriteFile(fullPath,nil,0644)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	tree.ToggleDir(tmpDir)

	assert.Equal(t,tree.IsMarked(tmpDir),true)

	tree.ToggleFile(filepath.Join(tmpDir,fileNames[0]))


	for _,name := range fileNames[1:] {
		testFilePath := filepath.Join(tmpDir,name)
		assert.Equal(t,tree.IsMarked(testFilePath),true)
	}

	assert.Equal(t,tree.IsMarked(filepath.Join(tmpDir,fileNames[0])),false)
} 


func TestInternalDropDir(t *testing.T) {
	tree := NewTree()
	tmpDir := t.TempDir()

	fileNames := []string{"a","b","c"}

	for _,name:= range fileNames{
		fullPath := filepath.Join(tmpDir,name)	
		err := os.Mkdir(fullPath,0644)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	tree.ToggleDir(tmpDir)

	assert.Equal(t,tree.IsMarked(tmpDir),true)

	tree.ToggleDir(filepath.Join(tmpDir,fileNames[0]))


	for _,name := range fileNames[1:] {
		testFilePath := filepath.Join(tmpDir,name)
		assert.Equal(t,tree.IsMarked(testFilePath),true)
	}

	assert.Equal(t,tree.IsMarked(filepath.Join(tmpDir,fileNames[0])),false)
} 

