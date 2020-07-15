package foo_bar

import "testing"

func TestFooBar(t *testing.T) {
	data := []int{
		3,
		5,
		15,
		1,
		4,
		113,
		45,
	}

	expected := []string {
		"foo",
		"bar",
		"15",
		"",
		"",
		"",
		"45",
	}

	for i, s := range data {
		if expected[i] != FooBar(s) {
			t.Errorf("ERROR: Expected: %v, got: %v", expected[i], FooBar(s))
		}
	}
}
