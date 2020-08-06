package aahw4

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (t *Tree) insert(data int) *Tree {
	if t.root == nil {
		t.root = &Node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (root *Node) height() int {
	var l int = 0
	var r int = 0

	if root == nil {
		//if reflect.DeepEqual(root, nil) {
		return 0
	}
	if root.left != nil {
		l = root.left.height() + 1
	}
	if root.right != nil {
		r = root.right.height() + 1
	}

	if r > l {
		return r
	} else {
		return l
	}
}

func (root *Node) lca(v1 int, v2 int) *Node {
	var ln *Node
	var rn *Node

	if root == nil {
		return nil
	}

	if (root.data == v1) || (root.data == v2) {
		return root
	}

	if root.left != nil {
		ln = root.left.lca(v1, v2)
	}
	if root.right != nil {
		rn = root.right.lca(v1, v2)
	}

	if (ln != nil) && (rn != nil) {
		return root
	}

	if ln != nil {
		return ln
	}
	if rn != nil {
		return rn
	}

	return nil
}

func (root *Node) bst() bool {
	if root == nil {
		return true
	}

	if (root.left != nil) && (root.data > root.left.data) {
		return false
	} else {
		return root.right.bst()
	}

	if (root.right != nil) && (root.data < root.right.data) {
		return false
	} else {
		return root.right.bst()
	}
}
