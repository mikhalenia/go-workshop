package dudkin

import "testing"


// пройден
func TestBalanced(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	} {
		{
			"{{}}{{",
			false,
		},
		{
			"{hello world!}",
			true,
		},
		{
			"()()()()()()",
			true,
		},
		{
			"[()]{}{[()()]()}",
			true,
		},
	}

	for i, test := range tests {
		res := Balanced(test.input)
		if res != test.expected {
			t.Errorf("Error on test %d: want %v, got %v", i, test.expected, res)
		}
	}

}
