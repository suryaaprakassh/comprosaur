package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/suryaaprakassh/comprosaur/marktree"
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
	Path string
	Kind FileKind
}

func NewFileType(name string,path string, isDir bool) FileType {
	if isDir {
		return FileType{
			Name: name,
			Path: path,
			Kind: Directory,
		}
	}

	return FileType{
		Name: name,
		Path: path,
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
	
	//tree to track mark status
	marktree *marktree.Tree
}

func (c *Cwd) moveForward() error {
	item, ok := c.Children.SelectedItem().(FileType)
	if !ok {
		return errors.New("Could Not Select Item!")
	}
	if item.Kind != Directory{
		return errors.New("The Item is Not a Directory!")
	}
	
	c.path = fmt.Sprintf("%s/%s",c.path,item.Name)

	return c.populateChildren()
}

func (c *Cwd) moveBack() error {
	index := strings.LastIndex(c.path,"/")
	if index == 0 {
			return errors.New("Cannot Move Back!")
	}
	c.path=c.path[:index]

	return c.populateChildren()
}

func (c *Cwd) markItem() error {
	index := c.Children.GlobalIndex()
	item, ok := c.Children.SelectedItem().(FileType)
	
	if item.Kind == File {
		c.marktree.ToggleFile(item.Path)
	} else {
		c.marktree.ToggleDir(item.Path)
	}

	if !ok {
		return errors.New("Could Not Select Item!")
	}
	c.Children.SetItem(index,item)
	return nil
}

func (c *Cwd) populateChildren() error {
	items := []list.Item{}
	files, err := os.ReadDir(c.path)
	if err != nil {
		return err
	}
	for _, child := range files {
		items = append(items, NewFileType(child.Name(), filepath.Join(c.path,child.Name()),child.IsDir()))
	}
	_ = c.Children.SetItems(items)
	return nil
}

func NewCwd() (*Cwd, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	marktree := marktree.NewTree()
	
	list := list.New(nil,itemDelegate{
		marktree: marktree,
	},20,14)
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
		marktree: marktree,
	}
	if err := c.populateChildren(); err != nil {
		return nil, err
	}
	return c, nil
}
