package prepare_exam

import "testing"

func TestSearch (t *testing.T) {
	type TestCase struct {
		data []int
		search int
		expected int
	}
	TestCases := []TestCase{
		{
			[]int{1, 3, 5, 6, 7},
			6,
			3,
		},
		{
			[]int{1, 3, 4, 10, 12},
			11,
			-1,
		},
	}

	for i, test := range TestCases {
		result := BinarySearch(test.data, test.search, 0, len(test.data) - 1)
		if result != test.expected {
			t.Errorf("Error: want %d, got %d, test case #%d", test.expected, result, i)
		}
	}
}

func DFS_test(t *testing.T) {
	cases := []struct {
		graph    map[string][]string
		search   string
		starting string
		expected bool
	}{
		{
			graph: map[string][]string {
				"a": {"b", "c"},
				"b": {"a", "c"},
				"c": {"a", "b", "d"},
				"d": {"c"},
			},
			search: "c",
			starting: "a",
			expected: true,
		},
		{
			graph: map[string][]string {
				"a": {"b", "c"},
				"b": {"a", "c"},
				"c": {"a", "b", "d"},
				"d": {"c"},
			},
			search: "e",
			starting: "a",
			expected: false,
		},
	}

	for i, el := range cases {
		v := make(map[string]bool)
		res := DFS(&el.graph, el.starting, el.search, &v)
		if res != el.expected {
			t.Errorf("ERROR on case %d: want %v, got %v", i, el.expected, res)
		}
	}
}
