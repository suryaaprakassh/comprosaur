package marktree

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/suryaaprakassh/comprosaur/context"
	"github.com/suryaaprakassh/comprosaur/stack"
)

func Dfs(node *Node, results *[]string, path string, is_dir bool) {
	if node.IsMarked() {
		if is_dir != node.is_dir {
			return
		}
		*results = append((*results), path)
		return
	}
	for p, child := range node.children {
		if !child.IsUnmark() {
			Dfs(child, results, filepath.Join(path, p), is_dir)
		}
	}
}

type Tree struct {
	root *Node

	ctx context.CTX
}

func (t *Tree) Clear() {
	t.root.Clear()
}

func (t *Tree) IsStatus(path string, status MarkedStatus) bool {
	n := t.root
	for key := range strings.SplitSeq(path, "/") {
		if key == "" {
			continue
		}
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
		if key == "" {
			continue
		}
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
		if key == "" {
			continue
		}
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

	//sets the current tracking count of the parent to the
	// number of elements in the directory
	if err := n.HandleParent(*parent, path); err != nil {
		return err
	}

	//drop the children if the entire directory gets marked
	if n.IsMarked() {
		n.current_count = n.item_count
		clear(n.children)
	}

	path = filepath.Dir(path)
	n = *parent
	s.Pop()

	//recursively set the status of the tree
	// for partial status
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
		if key == "" {
			continue
		}
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

// returns true if there is a dir marked
// if isRelative is false , then basepath is optional
// empty string can be passed
func (t *Tree) GetMarkedDirs(basepath string, isRelative bool) ([]string, bool) {
	paths := make([]string, 0)
	Dfs(t.root, &paths, "/", true)

	if isRelative {
		for idx, path := range paths {
			relPath, err := filepath.Rel(basepath, path)
			if err != nil {
				log.Fatal("Relative Path Convertsion Failed:", err.Error())
			}
			paths[idx] = relPath
		}
	}

	return paths, (len(paths) > 0)
}

// returns true if there is a file marked
// if isRelative is false , then basepath is optional
// empty string can be passed
func (t *Tree) GetMarkedFiles(basepath string,isRelative bool) ([]string, bool) {
	paths := make([]string, 0)
	Dfs(t.root, &paths, "/", false)

	if isRelative {
		for idx, path := range paths {
			relPath, err := filepath.Rel(basepath, path)
			if err != nil {
				log.Fatal("Relative Path Convertsion Failed:", err.Error())
			}
			paths[idx] = relPath
		}
	}

	return paths, (len(paths) > 0)
}

func NewTree(ctx context.CTX) *Tree {
	//Creating tree with root dir
	//TODO: change later
	t := &Tree{
		root: NewNode(true),
		ctx: ctx,
	}

	dirs, err := os.ReadDir("/")
	if err != nil {
		log.Fatal(err.Error())
	}
	t.root.item_count = len(dirs)
	return t
}
