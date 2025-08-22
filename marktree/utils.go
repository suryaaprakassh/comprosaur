package marktree

import (
	"path/filepath"
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
			Dfs(child, results, filepath.Join(path, p),is_dir)
		}
	}
}
