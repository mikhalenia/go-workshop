package savickaia

import "testing"

func TestMinNumber(t *testing.T) {
	type TestCaseFirst struct{
		data []int
		expected int
	}
	TestCases := []TestCaseFirst{
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{4, 15, 16, 18, 25},
			4,
		},
		{
			[]int{1, 2, 6, 40, 50},
			40,
		},
		{
			[]int{10, 39, 35, 43, 50},
			-1,
		},

	}
	for _, test := range TestCases {
		result := MinNumber(test.data)
		if result != test.expected{
			t.Errorf("Error: wanted : %d, get : %d", test.expected, result)
		}
	}
}

func TestAnagram(t *testing.T) {
	type TestCaseSecond struct{
		line string
		anagram string
		expected bool
	}
	TestCases := []TestCaseSecond{
		{
			"Hhelssddlo",
			"Hello",
			true,
		},
		{
			"",
			"Hello",
			false,
		},
		{
			"Hello programming !",
			"Hello world !",
			false,
		},
		{
			"cat",
			"act",
			true,
		},
	}
	for _, test := range TestCases {
		result := Anagram(test.line, test.anagram)
		if result != test.expected{
			t.Errorf("Error: wanted : %v, get : %v", test.expected, result)
		}
	}
}