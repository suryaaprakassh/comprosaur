package backend

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suryaaprakassh/comprosaur/shared"
	"github.com/suryaaprakassh/comprosaur/utils"
)

type MockSourceProvider struct {
	dirs  []string
	files []string
}

func (t *MockSourceProvider) GetMarkedDirs(basespath string, isRelative bool) ([]string, bool) {
	return t.dirs, len(t.dirs) > 0
}

func (t *MockSourceProvider) GetMarkedFiles(basespath string, isRelative bool) ([]string, bool) {
	return t.files, len(t.files) > 0
}

func MockNameProvider() string {
	return utils.RandString(10)
}

func NewTestSourceProvider() *MockSourceProvider {
	return &MockSourceProvider{
		dirs:  []string{"/foo/bar", "/foo/bar/baz"},
		files: []string{"/foo/bar/a.txt", "/foo/bar/b.txt"},
	}
}

func TestZip(t *testing.T) {
	logger := slog.Default()
	ctx := shared.New("/",logger)

	zip := NewZip(NewTestSourceProvider(),ctx)
	_, err := zip.Compress(false, MockNameProvider)
	assert.NoError(t, err)
	// err = cmd.Run()
	// assert.NoError(t, err)
}
