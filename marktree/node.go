package marktree

import (
	"os"
	"strings"
)

//TODO: optimise the unmarking
type Node struct {
	is_marked bool
	is_dir bool

	children  map[string]*Node
}

func (n *Node) ToggleMark() {
	n.is_marked = !n.is_marked
}

func (n *Node) Mark() {
	n.is_marked = true
}

func (n *Node) IsMarked() bool {
	return n.is_marked
}

func (n *Node) Repopulate(path string) error {
	index := strings.LastIndex(path,"/")
	parentPath := path[:index]
	fileName := path [index:]

	files, err := os.ReadDir(parentPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == fileName {
			continue
		}
		n.AddChild(file.Name(),file.IsDir())
	}

	return nil
}

func (n *Node) AddChild(name string, is_dir bool) {
	child := NewNode(is_dir)
	n.children[name] = child
} 

func NewNode(is_dir bool) *Node{
	return &Node{
		is_dir: is_dir,
		is_marked: false,
		children: make(map[string]*Node),
	}
}
