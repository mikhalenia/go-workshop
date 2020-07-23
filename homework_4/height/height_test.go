package height

import "testing"
import "studying_go_high_technologies_park/homework_4/binarytree"

type BSTTest struct {
	NodesValues []int
	Result int
}

func TestHeight(t *testing.T) {
	tests := []BSTTest{
		{[]int{3, 5, 2, 1, 4, 6, 7}, 3 },
		{[]int{15}, 0 },
		{[]int{3, 1, 7, 5, 4}, 3 },
	}
	for _, test := range tests {
		tree := binarytree.BinaryTree{}
		for _, elem := range test.NodesValues {
			tree.Insert(elem)
		}
		if test.Result != Height(tree.Root) {
			t.Errorf("ERROR: Expected: %v, got: %v", test.Result, Height(tree.Root))
		}
	}
}
