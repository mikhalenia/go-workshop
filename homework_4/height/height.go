package height

import "studying_go_high_technologies_park/homework_4/binarytree"

func Height(node *binarytree.BinaryNode) int {
	if node == nil {
		return -1
	}
	leftBranchHeight := Height(node.Left)
	rightBranchHeight := Height(node.Right)
	if leftBranchHeight >= rightBranchHeight {
		return leftBranchHeight + 1
	}
	return rightBranchHeight + 1
}


