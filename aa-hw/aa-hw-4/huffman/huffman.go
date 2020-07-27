package huffman

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	value rune
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

func parceTree(root *Node, array []int, i *int) rune {
	var j = *i

	if (root.left != nil) && (array[j] == '0') {
		*i++
		return parceTree(root.left, array, i)
	}
	if (root.right != nil) && (array[j] == '1') {
		*i++
		return parceTree(root.right, array, i)
	}
	return root.value
}

func (root *Node) huffman(array []int) string {
	var i int = 0

	var result []rune
	for i < len(array) {
		result = append(result, parceTree(root, array, &i))
	}
	return fmt.Sprint(result)
}
