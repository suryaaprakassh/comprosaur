package backend

import (
	"fmt"
	"os/exec"

	"github.com/suryaaprakassh/comprosaur/command"
)

type Zip struct {
	provider SourceProvider
}

func NewZip(provider SourceProvider) *Zip {
	return &Zip{
		provider: provider,
	}
}

func (c *Zip) Compress(verbose bool, np NameProvider) (*exec.Cmd, error) {
	dirs, haveDir := c.provider.GetMarkedDirs()
	files, haveFile := c.provider.GetMarkedFiles()

	if !haveFile && !haveDir {
		return nil, NoFileSelected
	}

	name := fmt.Sprintf("%s.zip",np())
	cmd := command.New("zip")
	cmd.Arg(name)

	if verbose {
		cmd.Arg("-v")
	}
	if haveDir {
		cmd.Arg("-r")
		cmd.Args(dirs...)
	}
	if haveFile {
		cmd.Args(files...)
	}
	return cmd.Build(), nil
}

func (c *Zip) EnsureInstallFatal() {

}


func (c *Zip) EnsureInstalled() error {
	return nil
}
