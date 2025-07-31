package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

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

	Marked bool
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
	if !ok {
		return errors.New("Could Not Select Item!")
	}
	item.Marked = !item.Marked
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
