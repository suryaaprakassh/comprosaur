package marktree

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestDir struct {
	path string
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
		path: tmp,
	}
}

func (t *TestDir) get(path string) string {
	return filepath.Join(t.path, path)
}

func TestTree(t *testing.T) {
	tree := NewTree()
	dirs := []string{
		"/foo/boo",
		"/foo/bar/baz",
	}
	files := []string{
		"/foo/bar/a.txt",
	}

	tDir := NewTestDir(t, dirs, files)

	for _, path := range dirs {
		err := tree.ToggleDir(tDir.get(path))
		assert.NoError(t, err)
	}

	for _, path := range dirs {
		p := tDir.get(path)
		assert.Equal(t, true, tree.IsMarked(p))
	}

	tree.ToggleFile(tDir.get(files[0]))

	for _, path := range dirs {
		tree.ToggleDir(tDir.get(path))
	}
	for _, path := range dirs {
		p := tDir.get(path)
		assert.Equal(t, false, tree.IsMarked(p))
	}
	assert.Equal(t, true, tree.IsMarked(tDir.get(files[0])))
}

func TestMarkDrop(t *testing.T) {
	tree := NewTree()
	dirs := []string{
		"/foo/bar/baz",
		"/foo/bar",
	}
	files := []string{
		"/foo/bar/a.txt",
	}
	tDir := NewTestDir(t, dirs, files)
	for _, path := range dirs {
		err := tree.ToggleDir(tDir.get(path))
		assert.NoError(t, err)
	}

	for _, path := range dirs {
		p := tDir.get(path)
		assert.Equal(t, true, tree.IsMarked(p))
	}

	tree.ToggleDir(tDir.get("/foo"))

	p := tDir.get(files[0])
	assert.Equal(t, true, tree.IsMarked(p))
}

func TestInternalDrop(t *testing.T) {
	tree := NewTree()

	files := []string{"a", "b", "c"}

	tDir := NewTestDir(t, nil, files)

	tree.ToggleDir(tDir.path)
	assert.Equal(t, true, tree.IsMarked(tDir.path))

	tree.ToggleFile(tDir.get(files[0]))

	for _, path := range files[1:] {
		p := tDir.get(path)
		assert.Equal(t, true, tree.IsMarked(p))
	}

	assert.Equal(t, false, tree.IsMarked(tDir.get(files[0])))
}

func TestInternalDropDir(t *testing.T) {
	tree := NewTree()

	dirs := []string{"a", "b", "c"}

	tDir := NewTestDir(t, dirs, nil)

	tree.ToggleDir(tDir.path)

	assert.Equal(t, true, tree.IsMarked(tDir.path))

	tree.ToggleDir(tDir.get(dirs[0]))

	for _, path := range dirs[1:] {
		p := tDir.get(path)
		assert.Equal(t, true, tree.IsMarked(p))
	}

	assert.Equal(t, false, tree.IsMarked(tDir.get(dirs[0])))
}

func TestMarkedReturn(t *testing.T) {
	tree := NewTree()
	dirs := []string{
		"/foo/boo",
		"/foo/bar",
	}
	files := []string{
		"/foo/bar/a.txt",
		"/foo/bar/b.txt",
	}
	tDir := NewTestDir(t, dirs, files)
	tree.ToggleDir(tDir.get(dirs[0]))
	tree.ToggleFile(tDir.get(files[0]))
	markedFiles, isMarked := tree.GetMarkedDirs()
	assert.True(t, isMarked)
	assert.Equal(t,1,len(markedFiles))
	assert.Equal(t, tDir.get(dirs[0]), markedFiles[0])

	markedFiles, isMarked = tree.GetMarkedFiles()
	assert.True(t, isMarked)
	assert.Equal(t,1,len(markedFiles))
	assert.Equal(t, tDir.get(files[0]), markedFiles[0])
}
