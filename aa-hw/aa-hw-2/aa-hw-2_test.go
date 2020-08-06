package aahw2

import (
	"reflect"
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

func TestRotLeft(t *testing.T) {

	type testValue struct {
		arr  []int32
		d    int32
		want []int32
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{
		[]int32{1, 2, 3, 4, 5},
		4,
		[]int32{5, 1, 2, 3, 4},
	})

	testValueArr = append(testValueArr, testValue{
		[]int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51},
		10,
		[]int32{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77},
	})

	testValueArr = append(testValueArr, testValue{
		[]int32{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97},
		13,
		[]int32{87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60},
	})

	for i, tVal := range testValueArr {
		result := rotLeft(tVal.arr, tVal.d)
		if !reflect.DeepEqual(tVal.want, result) {
			t.Errorf("Error Test hourglassSum. Iteration = %d, Want = %d and Got = '%v'.", i, tVal.want, result)
		}
	}
}

func TestAlternatingCharacters(t *testing.T) {

	type testValue struct {
		str  string
		want int32
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{"AAAA", 3})

	testValueArr = append(testValueArr, testValue{"BBBBB", 4})

	testValueArr = append(testValueArr, testValue{"ABABABAB", 0})

	testValueArr = append(testValueArr, testValue{"BABABA", 0})

	testValueArr = append(testValueArr, testValue{"AAABBB", 4})

	for i, tVal := range testValueArr {
		result := alternatingCharacters(tVal.str)
		if tVal.want != result {
			t.Errorf("Error Test hourglassSum. Iteration = %d, Want = %d and Got = '%v'.", i, tVal.want, result)
		}
	}
}
