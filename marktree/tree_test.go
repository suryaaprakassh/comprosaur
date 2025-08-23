package marktree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suryaaprakassh/comprosaur/utils"
)


func TestTree(t *testing.T) {
	tree := NewTree()
	dirs := []string{
		"/foo/boo",
		"/foo/bar/baz",
	}
	files := []string{
		"/foo/bar/a.txt",
	}

	tDir := utils.NewTestDir(t, dirs, files)

	for _, path := range dirs {
		err := tree.ToggleDir(tDir.Get(path))
		assert.NoError(t, err)
	}

	for _, path := range dirs {
		p := tDir.Get(path)
		assert.Equal(t, true, tree.IsMarked(p))
	}

	tree.ToggleFile(tDir.Get(files[0]))

	for _, path := range dirs {
		tree.ToggleDir(tDir.Get(path))
	}
	for _, path := range dirs {
		p := tDir.Get(path)
		assert.Equal(t, false, tree.IsMarked(p))
	}
	assert.Equal(t, true, tree.IsMarked(tDir.Get(files[0])))
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
	tDir := utils.NewTestDir(t, dirs, files)
	for _, path := range dirs {
		err := tree.ToggleDir(tDir.Get(path))
		assert.NoError(t, err)
	}

	for _, path := range dirs {
		p := tDir.Get(path)
		assert.Equal(t, true, tree.IsMarked(p))
	}

	tree.ToggleDir(tDir.Get("/foo"))

	p := tDir.Get(files[0])
	assert.Equal(t, true, tree.IsMarked(p))
}

func TestInternalDrop(t *testing.T) {
	tree := NewTree()

	files := []string{"a", "b", "c"}

	tDir := utils.NewTestDir(t, nil, files)

	tree.ToggleDir(tDir.Path)
	assert.Equal(t, true, tree.IsMarked(tDir.Path))

	tree.ToggleFile(tDir.Get(files[0]))

	for _, path := range files[1:] {
		p := tDir.Get(path)
		assert.Equal(t, true, tree.IsMarked(p))
	}

	assert.Equal(t, false, tree.IsMarked(tDir.Get(files[0])))
}

func TestInternalDropDir(t *testing.T) {
	tree := NewTree()

	dirs := []string{"a", "b", "c"}

	tDir := utils.NewTestDir(t, dirs, nil)

	tree.ToggleDir(tDir.Path)

	assert.Equal(t, true, tree.IsMarked(tDir.Path))

	tree.ToggleDir(tDir.Get(dirs[0]))

	for _, path := range dirs[1:] {
		p := tDir.Get(path)
		assert.Equal(t, true, tree.IsMarked(p))
	}

	assert.Equal(t, false, tree.IsMarked(tDir.Get(dirs[0])))
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
	tDir := utils.NewTestDir(t, dirs, files)
	tree.ToggleDir(tDir.Get(dirs[0]))
	tree.ToggleFile(tDir.Get(files[0]))
	markedFiles, isMarked := tree.GetMarkedDirs()
	assert.True(t, isMarked)
	assert.Equal(t,1,len(markedFiles))
	assert.Equal(t, tDir.Get(dirs[0]), markedFiles[0])

	markedFiles, isMarked = tree.GetMarkedFiles()
	assert.True(t, isMarked)
	assert.Equal(t,1,len(markedFiles))
	assert.Equal(t, tDir.Get(files[0]), markedFiles[0])
}
