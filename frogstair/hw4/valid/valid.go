package valid

// Node represents a node in a tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

var values map[int]bool

func init() {
	values = make(map[int]bool)
}

// IsValid checks if a tree is valid
func (n *Node) IsValid() bool {
	if n == nil {
		return true
	}

	res := true

	if n.Left != nil && n.Left.Value > n.Value {
		res = false
	}
	if n.Right != nil && n.Right.Value < n.Value {
		res = false
	}

	_, ok := values[n.Value]
	if ok {
		res = false
	} else {
		values[n.Value] = true
	}

	return res && n.Left.IsValid() && n.Right.IsValid()
}
