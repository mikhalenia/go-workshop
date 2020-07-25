package mglinnik

import "testing"

// TestFoobar ...
func TestFoobar(t *testing.T) {
	input := []int{3, 5, 15}
	expected := []string{"foo", "bar", "15"}

	for i, x := range input {
		if Foobar(x) != expected[i] {
			t.Errorf("Expected %d to equal %s, but got %s", input[i], expected[i], Foobar(x))
		}
	}
}
