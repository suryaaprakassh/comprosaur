package backend

import (
	"fmt"
	"os/exec"

	"github.com/suryaaprakassh/comprosaur/command"
	"github.com/suryaaprakassh/comprosaur/context"
	"github.com/suryaaprakassh/comprosaur/utils"
)

type Zip struct {
	provider SourceProvider

	ctx context.CTX
}

func NewZip(provider SourceProvider, ctx context.CTX) *Zip {
	return &Zip{
		provider: provider,
		ctx:      ctx,
	}
}

func (c *Zip) Compress(verbose bool, np NameProvider) (*exec.Cmd, error) {
	dirs, haveDir := c.provider.GetMarkedDirs(c.ctx.GetPath(), true)
	files, haveFile := c.provider.GetMarkedFiles(c.ctx.GetPath(), true)

	if !haveFile && !haveDir {
		return nil, NoFileSelected
	}

	name := fmt.Sprintf("%s.zip", np())
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
	if !utils.IsInstalled("zip") {
		panic("Zip is not installed!")
	}
}

func (c *Zip) EnsureInstalled() error {
	if !utils.IsInstalled("zip") {
		return fmt.Errorf("Zip is not installed!")
	}
	return nil
}
