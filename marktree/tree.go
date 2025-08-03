package marktree

import (
	"log"
	"strings"
)

type Tree struct {
	root *Node
}

func (t *Tree) IsMarked(path string) bool {
	n := t.root
	for key := range strings.SplitSeq(path, "/") {
		node, ok := n.children[key]
		if ok {
			n = node
			if n.is_dir && n.is_marked {
				return true
			}
		} else {
			return false
		}
	}
	return n.IsMarked()
}

func (t *Tree) ToggleDir(path string) error {
	var parent *Node 
	n := t.root
	for key := range strings.SplitSeq(path, "/") {
		node, ok := n.children[key]
		if ok {
			parent = n
			n = node
		} else {
			n.AddChild(key, true)
			n, _ = n.children[key]
		}
	}
	n.ToggleMark()
	if n.IsMarked() {
		clear(n.children)
	}

	if !n.IsMarked() && parent.IsMarked() {
		parent.is_marked = false
		return parent.Repopulate(path)
	}

	return nil
}

func (t *Tree) ToggleFile(path string) error {
	var parent *Node
	n := t.root
	for key := range strings.SplitSeq(path, "/") {
		node, ok := n.children[key]
		if ok {
			parent = n
			n = node
		} else {
			n.AddChild(key, true)
			n, _ = n.children[key]
		}
	}
	if parent.IsMarked() {
			n.is_marked = false
	}else{
			n.is_marked = true 
	}
	//making it a file here
	n.is_dir = false
	log.Println(path,n.IsMarked(),"parent=",parent.IsMarked())
	if !n.IsMarked() && parent.IsMarked() {
		parent.is_marked = false
		return parent.Repopulate(path)
	}
	log.Println(path,n.IsMarked(),"parent=",parent.IsMarked())
	return nil
}

func NewTree() *Tree {
	//Creating tree with root dir
	//TODO: change later
	return &Tree{
		root: NewNode(true),
	}
}
