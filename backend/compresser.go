package backend

import (
	"errors"
	"os/exec"
)

// function that provides name for the compressed
type NameProvider func() string

var NoFileSelected = errors.New("No File Selected!")
var DirectoriesSelected = errors.New("Directories cannot be extracted!")
var MoreFilesSelected = errors.New("More than 1 file is selected!")

type Compresser interface {
	Compress(verbose bool, np NameProvider) (*exec.Cmd, error)
	EnsureInstalled() error

	//should crash if not installed
	EnsureInstallFatal()
}

type Extractor interface {
	Extract(verbose bool,np NameProvider) (*exec.Cmd,error)


	EnsureInstalled() error
	//should crash if not installed
	EnsureInstallFatal()

}

type Backend interface {
	Compresser
	Extractor
}

type SourceProvider interface {
	// gets the relative path of the dirs from the current dir
	// basepath - base path of the current dir
	// isRelative - if  relative path is needed
	//         base path is optional if its not relative
	GetMarkedDirs(basepath string, isRelative bool) ([]string, bool)
	// gets the relative path of the dirs from the current dir
	// basepath - base path of the current dir
	// isRelative- if  relative path is needed
	//         base path is optional if its not relative
	GetMarkedFiles(basepath string, isRelative bool) ([]string, bool)
}
