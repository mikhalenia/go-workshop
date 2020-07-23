package binarytree

type BinaryNode struct {
	Data  int
	Left  *BinaryNode
	Right *BinaryNode
}

type BinaryTree struct {
	Root *BinaryNode
}

func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{Data: data, Left: nil, Right: nil}
	} else {
		t.Root.Insert(data)
	}
	return t
}

func (node *BinaryNode) Insert(data int) {
	if node == nil {
		return
	} else if data <= node.Data {
		if node.Left == nil {
			node.Left = &BinaryNode{Data: data, Left: nil, Right: nil}
		} else {
			node.Left.Insert(data)
		}
	} else {
		if node.Right == nil {
			node.Right = &BinaryNode{Data: data, Left: nil, Right: nil}
		} else {
			node.Right.Insert(data)
		}
	}
}
