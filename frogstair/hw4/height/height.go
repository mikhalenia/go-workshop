package height

// Node represents a node in a tree
type Node struct {
	Left  *Node
	Right *Node
}

// Height finds the height of a trees
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return max(n.Left.Height(), n.Left.Height()) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	} else {
		return a
	}
}
