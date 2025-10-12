package backend

import (
	"fmt"
	"os/exec"

	"github.com/suryaaprakassh/comprosaur/command"
	"github.com/suryaaprakassh/comprosaur/shared"
	"github.com/suryaaprakassh/comprosaur/utils"
)

type Zip struct {
	provider SourceProvider

	ctx shared.CTX
}

func NewZip(provider SourceProvider, ctx shared.CTX) *Zip {
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
	
	//TODO: split the name provider and take the path from the ctx
	name := fmt.Sprintf("%s/%s.zip",c.ctx.GetPath(), np())
	cmd := command.New("zip")
	cmd.Arg(name)

	if verbose {
		cmd.Arg("-o")
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


func (c *Zip) Extract(verbose bool, np NameProvider) (*exec.Cmd, error) {
	_ , haveDir := c.provider.GetMarkedDirs(c.ctx.GetPath(), true)
	files, haveFile := c.provider.GetMarkedFiles(c.ctx.GetPath(), true)

	if haveDir {
		return nil, DirectoriesSelected
	}

	if !haveFile {
		return nil , NoFileSelected
	}

	if len(files) > 1 { 
		return nil, MoreFilesSelected
	}
	path := fmt.Sprintf("%s/%s",c.ctx.GetPath(),np())
	cmd := command.New("unzip")
	cmd.Arg(files[0])
	cmd.Arg("-d")
	cmd.Arg(path)

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
