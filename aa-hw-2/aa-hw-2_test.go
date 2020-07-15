package aahw2

import (
	"testing"
)

func TestHourglassSum(t *testing.T) {

	type testValue struct {
		matrix [][]int32
		want   int32
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{
		[][]int32{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
		28,
	})

	testValueArr = append(testValueArr, testValue{
		[][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 9, 2, -4, -4, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, -1, -2, -4, 0},
		},
		13,
	})

	testValueArr = append(testValueArr, testValue{
		[][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
		19,
	})

	for i, tVal := range testValueArr {
		result := hourglassSum(tVal.matrix)
		if tVal.want != result {
			t.Errorf("Error Test hourglassSum. Iteration = %d, Want = %d and Got = '%v'.", i, tVal.want, result)
		}

	}
}
