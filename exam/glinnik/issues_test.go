package glinnik

import "testing"

func TestPalindrome (t *testing.T) {
	type TestCase struct {
		input string
		expected bool
}
	TestCases := []TestCase{
	{
		"abba",
		true,
	},
	{
		"ahah",
		false,
	},
}

	for i, test := range TestCases {
		result := Palindrome(test.input)
		if result != test.expected {
			t.Errorf ("Error want %v, got %v, testid %d", test.expected, result, i)
		}
	}
}

func TestMinGraphElement(t *testing.T) {
	cases := []struct {
		input    map[int]int
		expected int
	}{
		{
		input: map[int]int{
				1: 2,
				2: 3,
				3: 4,
			},
		expected: 1,
		},
		{
		input:
			map[int]int{
				2: 1,
				8: 4,
				3: 3,
				4: 6,
				5: 7,
			},
		expected: 2,
		},
	}

	for i, el := range cases {
		result := MinGraphElement(el.input)
		if result != el.expected {
			t.Errorf("Error: want %d, got %d, testid %d", el.expected, result, i)
		}
	}
}
