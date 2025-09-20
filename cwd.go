package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/suryaaprakassh/comprosaur/backend"
	"github.com/suryaaprakassh/comprosaur/context"
	"github.com/suryaaprakassh/comprosaur/marktree"
	"github.com/suryaaprakassh/comprosaur/utils"
)

type Cwd struct {
	Children list.Model

	//tree to track mark status
	marktree *marktree.Tree
	backend  backend.Backend

	ctx context.CTX

	selectedItem int
}

func (c *Cwd) moveForward() error {
	c.selectedItem = c.Children.GlobalIndex()

	item, ok := c.Children.SelectedItem().(FileType)

	if !ok {
		return errors.New("Could Not Select Item!")
	}
	if item.Kind != Directory {
		return errors.New("The Item is Not a Directory!")
	}
	c.ctx.AppendPath(item.Name)
	return c.populateChildren()
}

func (c *Cwd) moveBack() error {
	path := c.ctx.GetPath()

	index := strings.LastIndex(path, "/")
	if index == 0 {
		return errors.New("Cannot Move Back!")
	}
	c.ctx.UpdatePath(path[:index])

	return c.populateChildren()
}

func (c *Cwd) compressSelected() error {

	np := func() string {
		return filepath.Join(c.ctx.GetPath(), utils.RandString(5))
	}
	//TODO: have a state for verbose
	//default set to true
	cmd, err := c.backend.Compress(true, np)
	if err != nil {
		return err
	}

	//TODO: this blocks do something about it
	if err := cmd.Run(); err != nil {
		return err
	}

	c.ctx.UpdateStatus("Compression Successful!")
	//redraw ui
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
	c.Children.SetItem(index, item)
	return nil
}

func (c *Cwd) populateChildren() error {
	items := []list.Item{}
	files, err := os.ReadDir(c.ctx.GetPath())
	if err != nil {
		return err
	}
	for _, child := range files {
		items = append(items, NewFileType(child.Name(), filepath.Join(c.ctx.GetPath(), child.Name()), child.IsDir()))
	}
	_ = c.Children.SetItems(items)

	c.Children.Select(c.selectedItem)
	return nil
}

func NewCwd(ctx context.CTX) (*Cwd, error) {
	marktree := marktree.NewTree()
	list := list.New(nil, itemDelegate{
		marktree: marktree,
	}, 20, 14)
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)
	list.SetShowTitle(false)
	list.Styles.Title = titleStyle
	list.Styles.PaginationStyle = paginationStyle
	list.Styles.HelpStyle = helpStyle
	c := &Cwd{
		Children: list,
		marktree: marktree,

		backend: backend.NewZip(marktree, ctx),
		ctx:     ctx,

		selectedItem: 0,
	}
	if err := c.populateChildren(); err != nil {
		return nil, err
	}
	return c, nil
}
