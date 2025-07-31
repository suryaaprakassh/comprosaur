package main

import (
	"os"
	"github.com/charmbracelet/bubbles/list"
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

// to implement bubble tea list.Item 
func (f FileType) FilterValue() string {
	return f.Name
}

func (f *FileType) String() string {
	return f.Name + f.Kind.String()
}

type Cwd struct {
	path     string
	Children list.Model
	length   int 
}

func (c *Cwd) populateChildren() error {
	items := []list.Item{}
	files, err := os.ReadDir(c.path)
	if err != nil {
		return err
	}
	for _, child := range files {
		items = append(items, NewFileType(child.Name(), child.IsDir()))
	}
	_ = c.Children.SetItems(items)
	return nil
}

func NewCwd() (*Cwd, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	
	list := list.New(nil,itemDelegate{},20,14)
	list.SetShowStatusBar(false)	
	list.SetFilteringEnabled(false)
	list.SetShowTitle(false)
	list.Styles.Title=titleStyle
	list.Styles.PaginationStyle = paginationStyle
	list.Styles.HelpStyle = helpStyle
	c := &Cwd{
		path: path,
		length: 0,
		Children: list,
	}
	if err := c.populateChildren(); err != nil {
		return nil, err
	}
	return c, nil
}
