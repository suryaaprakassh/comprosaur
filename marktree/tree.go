package marktree

import (
	"strings"
)

type Tree struct {
	root *Node
}

func (t *Tree) IsStatus(path string,status MarkedStatus) bool {
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
			if n.is_dir && n.IsMarked(){
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
			parent = n
			n.AddChild(key, true)
			n, _ = n.children[key]
		}
	}
	
	if err := n.HandleParent(parent,path); err != nil {
		return err
	}
	
	//drop the children if the entire directory gets marked
	if n.IsMarked() {
		clear(n.children)
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
			parent = n 
			n.AddChild(key, true)
			n, _ = n.children[key]
		}
	}
	//making it a file here
	n.is_dir = false

	return n.HandleParent(parent,path)
}

func NewTree() *Tree {
	//Creating tree with root dir
	//TODO: change later
	return &Tree{
		root: NewNode(true),
	}
}
