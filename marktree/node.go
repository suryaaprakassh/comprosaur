package marktree

import (
	"os"
	"path/filepath"
	"strings"

	"log"
)

type MarkedStatus int

const (
	Unmarked MarkedStatus = iota
	Marked
	Partial
)

//impl of stringer interface for debug purposes
func (n MarkedStatus) String() string {
	switch n {
	case Unmarked:
			return "Unmarked"
	case Marked:
			return "Marked"
	case Partial:
			return "Partial"
	}
	panic("Unreachable!")
}

// TODO: optimise the unmarking
type Node struct {
	status MarkedStatus
	is_dir bool

	children      map[string]*Node
	item_count    int
	current_count int
}


func (n *Node) Mark(status MarkedStatus) {
	n.status = status
}

func (n *Node) IsMarked() bool {
	return n.status == Marked
}

func (n *Node) IsPartial() bool {
	return n.status == Partial
}

func (n *Node) IsUnmark() bool {
	return n.status == Unmarked
}

func (n *Node) HandleRetriggerStatus(parent *Node) {
	if n.IsMarked() {
		parent.current_count += 1
	} else if n.IsUnmark() {
		parent.current_count -= 1
	}

	if n.IsPartial() {
		parent.Mark(Partial)
		return
	}

	if parent.item_count == parent.current_count {
		parent.Mark(Marked)
	} else if parent.current_count == 0 {
		parent.Mark(Unmarked)
	} else {
		parent.Mark(Partial)
	}
}

func (n *Node) HandleParent(parent *Node, path string) error {
	if parent.IsMarked() {
		parent.current_count -= 1
		n.Mark(Unmarked)
		if err := parent.Repopulate(path); err != nil {
			return err
		}
	} else {
		if n.IsMarked() {
			parent.current_count -= 1
			n.Mark(Unmarked)
		} else {
			parent.current_count += 1
			n.Mark(Marked)
		}
	}
	switch parent.current_count {
	case parent.item_count:
		parent.Mark(Marked)
	case 0:
		parent.Mark(Unmarked)
	default:
		parent.Mark(Partial)
	}
	return nil
}

func (n *Node) Repopulate(path string) error {
	index := strings.LastIndex(path, "/")
	parentPath := path[:index]
	fileName := path[index+1:]

	files, err := os.ReadDir(parentPath)

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == fileName {
			continue
		}

		childPath := filepath.Join(parentPath, file.Name())
		child := n.AddChild(file.Name(), file.IsDir(), childPath)
		child.status = Marked
	}
	
	n.current_count= n.item_count - 1

	return nil
}

func (n *Node) AddChild(name string, is_dir bool, path string) *Node {
	child := NewNode(is_dir)
	n.children[name] = child

	//callibrating the count of the children for partial status
	if is_dir {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal("Path does not exist",path)
		}
		child.item_count = len(files)
	}
	return child
}

func NewNode(is_dir bool) *Node {
	return &Node{
		is_dir:        is_dir,
		status:        Unmarked,
		children:      make(map[string]*Node),
		item_count:    0,
		current_count: 0,
	}
}
