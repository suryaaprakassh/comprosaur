package marktree

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/suryaaprakassh/comprosaur/stack"
)

type Tree struct {
	root *Node
}

func (t *Tree) IsStatus(path string, status MarkedStatus) bool {
	n := t.root
	for key := range strings.SplitSeq(path, "/") {
		node, ok := n.children[key]
		if ok {
			n = node
		} else {
			return false
		}
	}
	return n.status == status
}

func (t *Tree) IsMarked(path string) bool {
	n := t.root
	for key := range strings.SplitSeq(path, "/") {
		node, ok := n.children[key]
		if ok {
			n = node
			if n.is_dir && n.IsMarked() {
				return true
			}
		} else {
			return false
		}
	}
	return n.IsMarked()
}

func (t *Tree) ToggleDir(path string) error {
	n := t.root
	s := stack.New[*Node]()

	currPath := "/"
	for key := range strings.SplitSeq(path, "/") {
		node, ok := n.children[key]
		currPath = filepath.Join(currPath, key)

		if ok {
			s.Push(n)
			n = node
		} else {
			s.Push(n)
			n.AddChild(key, true, currPath)
			n, _ = n.children[key]
		}
	}

	parent, err := s.Top()
	if err != nil {
		return err
	}
	if err := n.HandleParent(*parent, path); err != nil {
		return err
	}
	path = filepath.Dir(path)
	n = *parent
	s.Pop()
	//drop the children if the entire directory gets marked
	if n.IsMarked() {
		clear(n.children)
	}

	for !s.IsEmpty() {
		parent, err := s.Top()
		if err != nil {
			return err
		}
		n.HandleRetriggerStatus(*parent)
		n = *parent
		s.Pop()
	}

	return nil
}

func (t *Tree) ToggleFile(path string) error {
	n := t.root

	s := stack.New[*Node]()

	parts := strings.Split(path, "/")
	currPath := "/"

	for idx, key := range parts {
		node, ok := n.children[key]
		currPath = filepath.Join(currPath, key)

		if ok {
			s.Push(n)
			n = node
		} else {
			s.Push(n)
			n.AddChild(key, idx+1 != len(parts), currPath)
			n, _ = n.children[key]
		}
	}

	parent, err := s.Top()
	if err != nil {
		return err
	}
	if err := n.HandleParent(*parent, path); err != nil {
		return err
	}
	path = filepath.Dir(path)
	n = *parent
	s.Pop()

	for !s.IsEmpty() {
		parent, err := s.Top()
		if err != nil {
			return err
		}
		n.HandleRetriggerStatus(*parent)
		n = *parent
		s.Pop()
	}

	return nil
}

func NewTree() *Tree {
	//Creating tree with root dir
	//TODO: change later
	t := &Tree{
		root: NewNode(true),
	}

	dirs, err := os.ReadDir("/")
	if err != nil {
		log.Fatal(err.Error())
	}
	t.root.item_count = len(dirs)
	return t
}
