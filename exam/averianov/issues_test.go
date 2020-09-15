package averianov

import (
	"testing"
)

func TestLeastInArray(t *testing.T) {
	type ChecketValues struct {
		arr      []int
		expected int
	}

	var checkArr []ChecketValues

	checkArr = append(checkArr, ChecketValues{[]int{1, 2, 5, 9, 7, 45, 156, 789}, 9})
	checkArr = append(checkArr, ChecketValues{[]int{756, 65, 33, 15, 7, 6, 2, 1}, 6})

	for i, tested := range checkArr {
		result := leastInArray(tested.arr)
		if result != tested.expected {
			t.Errorf("Error Test leastInArray. Iteration = %d, Want = %d and Got = '%v'.", i, tested.expected, result)
		}
	}

}

func TestIterateRuns(t *testing.T) {

	type ChecketValues struct {
		str      string
		value    []rune
		expected map[rune]int
	}

	var checkArr []ChecketValues

	checkArr = append(checkArr, ChecketValues{"AndOdddPoSd", []rune{'A', 'd'}, map[rune]int{'A': 1, 'd': 5}})

	for _, tested := range checkArr {
		result := iterateRuns(tested.str, tested.value)

		for key, value := range result {
			expectedValue := tested.expected[key]
			if expectedValue != value {
				t.Errorf("Error iterateRuns key: %#U, value: %d, expected: %d", key, value, expectedValue)
			} else {
				t.Logf("%#U: %d", key, value)
			}
		}
	}

}
