package marktree

import (
	"fmt"
	"strings"
)

func (t *Tree) DebugPath(path string){
	n := t.root
	sequence := strings.Split(path,"/")
	for idx,key := range sequence {
		if key == "" {
			continue
		}
		node, ok := n.children[key]
		if ok {
			n = node
			fmt.Printf("(%s,%s)",key,node.status)
		} else { 
			fmt.Printf("(%s,NO)",key)
		}

		if idx + 1 < len(sequence) {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

