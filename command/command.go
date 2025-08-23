package command 

import "os/exec"

type CommandBuilder struct{
	name string
	args []string
}

func New(name string) *CommandBuilder {
	return &CommandBuilder{
		name: name,
		args: make([]string,0),
	}
}

func (c *CommandBuilder) Arg(arg string) *CommandBuilder {
	c.args = append(c.args, arg)
	return c
}


func (c *CommandBuilder) Args(arg ...string) *CommandBuilder {
	c.args = append(c.args, arg...)
	return c
}

func (c *CommandBuilder) Build() *exec.Cmd {
	return exec.Command(c.name,c.args...)
}
