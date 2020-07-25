package lca

import "studying_go_high_technologies_park/homework_4/binarytree"

func LowestCommonAncestor(node *binarytree.BinaryNode, v1, v2 int) *binarytree.BinaryNode {
	if node == nil {
		return nil
	}
	if node.Data > v1 && node.Data > v2 {
		return LowestCommonAncestor(node.Left, v1, v2)
	}
	if node.Data < v1 && node.Data < v2 {
		return LowestCommonAncestor(node.Right, v1, v2)
	}
	return node
}
