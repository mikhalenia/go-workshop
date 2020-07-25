package valid

import (
	"testing"
)

// TestValid tests the height of the nodes
func TestValid(t *testing.T) {
	type test struct {
		TestNode *Node
		Valid    bool
	}

	data := []test{
		{
			&Node{
				10,
				&Node{
					5,
					&Node{
						4,
						nil,
						nil,
					},
					nil,
				},
				&Node{
					11,
					&Node{
						8,
						nil,
						nil,
					},
					&Node{
						15,
						nil,
						nil,
					},
				},
			},
			true,
		},
		{
			&Node{
				3,
				&Node{
					2,
					&Node{
						1,
						nil,
						nil,
					},
					nil,
				},
				&Node{
					5,
					&Node{
						6,
						nil,
						nil,
					},
					&Node{
						1,
						nil,
						nil,
					},
				},
			},
			false,
		},
	}

	for i, el := range data {
		if el.Valid != el.TestNode.IsValid() {
			t.Errorf("Error test case %d: want %t, got %t", i, el.Valid, el.TestNode.IsValid())
		}
	}
}
