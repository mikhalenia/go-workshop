package height

import (
	"testing"
)

// TestHeight tests the height of the nodes
func TestHeight(t *testing.T) {
	type test struct {
		TestNode *Node
		Height   int
	}

	data := []test{
		{
			TestNode: &Node{
				&Node{
					&Node{nil, nil},
					&Node{nil, nil},
				},
				&Node{
					&Node{
						&Node{nil, nil},
						nil,
					},
					&Node{nil, nil},
				},
			},
			Height: 3,
		},
	}

	for _, el := range data {
		h := el.TestNode.Height()
		if h != el.Height {
			t.Errorf("Error: want %d, got %d", el.Height, h)
		}
	}

}
