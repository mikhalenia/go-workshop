package my_package

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	type TestCase struct {
		line string
		secret string
		expected bool
	}
	testCases := []TestCase{
		//{"I love the Beatles", "I love you", false},
		{"programming", "prog", true},
	}
	for i, testCase := range testCases {
		if testCase.expected != IsAnagram(testCase.line, testCase.secret){
			t.Errorf("ERROR: Expected: %v, case numb: %v", testCase.expected, i)
		}
	}

}
