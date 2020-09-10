package dudkin

import "testing"

/*
 *     15
      /  \
     5    7
    /    / \
   3    1   6
 */


func TestSmallest(t *testing.T) {
	tests := []struct {
		node *Node
		expected int
	}{
		{
			&Node{
				15,
				&Node{
					5,
					&Node{
						3,
						nil,
						nil,
					},
					nil,
				},
				&Node{
					7,
					&Node{
						1,
						nil,
						nil,
					},
					&Node{
						6,
						nil,
						nil,
					},
				},
			},
			1,
		},
	}

	for _, test := range tests {
		e := test.node.FindSmallest(test.node.v)
		if e != test.expected {
			t.Errorf("Error, want %d, got %d", test.expected, e)
		}
	}
}