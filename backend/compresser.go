package backend

import (
	"errors"
	"os/exec"
)

// function that provides name for the compressed
type NameProvider func() string

var NoFileSelected = errors.New("No File Selected!")

type Compresser interface {
	Compress(verbose bool, np NameProvider) (*exec.Cmd, error)
	EnsureInstalled() error

	//should crash if not installed
	EnsureInstallFatal()
}

type Extractor interface {
}

type Backend interface {
	Compresser
	Extractor
}

type SourceProvider interface {
	GetMarkedDirs() ([]string, bool)
	GetMarkedFiles() ([]string, bool)
}
