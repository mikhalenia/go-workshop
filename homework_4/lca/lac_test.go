package lca

import "testing"
import "studying_go_high_technologies_park/homework_4/binarytree"

type BSTTest struct {
	nodesValues []int
	v1, v2, result int
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []BSTTest{
		{[]int{4, 2, 3, 1, 7, 6},1, 7, 4 },
	}
	for _, test := range tests {
		tree := &binarytree.BinaryTree{}
			for _, elem := range test.nodesValues {
			tree.Insert(elem)
		}
			if test.result != LowestCommonAncestor(tree.Root, test.v1, test.v2).Data {
			t.Errorf("ERROR: Expected: %v, got: %v", test.result, LowestCommonAncestor(tree.Root, test.v1, test.v2))
		}

	}
}
