package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type FileKind int

func (f FileKind) String() string {
	switch f {
	case Directory:
		return "/"
	}
	return ""
}

const (
	Directory FileKind = iota
	File
)

type FileType struct {
	Name string
	Kind FileKind
}

func NewFileType(name string, isDir bool) FileType {
	if isDir {
		return FileType{
			Name: name,
			Kind: Directory,
		}
	}

	return FileType{
		Name: name,
		Kind: File,
	}
}

func (f *FileType) String() string {
	return f.Name + f.Kind.String()
}

type Cwd struct {
	path     string
	Children []FileType
}

func (c *Cwd) populateChildren() error {
	c.Children = nil
	files,err := os.ReadDir(c.path)
	if err != nil {
		return err
	}
	for _,child := range files  {
		c.Children = append(c.Children, NewFileType(child.Name(),child.IsDir()))	
	}
	return nil
}

func (c *Cwd) moveBack() error  {
	idx := strings.LastIndex(c.path,"/")
	if idx != -1 {
		c.path = c.path[:idx]
		return c.populateChildren()
	}
	return errors.New("Could Not Move Back!")
}

func (c *Cwd) selectChild(idx int) error {
	child := c.Children[idx]
	switch child.Kind {
	case Directory:
			c.path = fmt.Sprintf("%s/%s",c.path,child.Name)	
			return c.populateChildren()
	}
	return nil
}

func NewCwd() (*Cwd,error){
	path , err := os.Getwd()
	if err != nil {
		return nil,err
	}
	c := &Cwd{
		path: path,
	}
	if err := c.populateChildren(); err != nil {
		return nil,err
	}
	return c,nil
}
