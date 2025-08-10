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

//TODO: optimise the unmarking
type Node struct {
	status MarkedStatus	
	is_dir bool

	children  map[string]*Node
	item_count int
}

func (n *Node) Mark(status MarkedStatus) {
	n.status= status
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

func (n *Node) HandleParent(parent *Node,path string) error {
		if parent.IsMarked() {
			parent.Mark(Partial)
			n.Mark(Unmarked)
			return parent.Repopulate(path)
		}else {
			if n. IsMarked() {
				n.Mark(Unmarked)
				parent.Mark(Unmarked)
			}else {
				parent.Mark(Partial)
				n.Mark(Marked)
			}
		}
		return nil
}

func (n *Node) Repopulate(path string) error {
	index := strings.LastIndex(path,"/")
	parentPath := path[:index]
	fileName := path [index+1:]

	files, err := os.ReadDir(parentPath)

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == fileName {
			continue
		}

		childPath := filepath.Join(parentPath,file.Name())
		child := n.AddChild(file.Name(),file.IsDir(),childPath)

		child.status = Marked
	}

	return nil
}

func (n *Node) AddChild(name string, is_dir bool,path string) *Node {
	child := NewNode(is_dir)
	n.children[name] = child
	
	//callibrating the count of the children for partial status 
	if is_dir {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(path)
		}
		child.item_count = len(files)
	}

	return child
} 

func NewNode(is_dir bool) *Node{
	return &Node{
		is_dir: is_dir,
		status: Unmarked,
		children: make(map[string]*Node),
		item_count: 0,
	}
}
