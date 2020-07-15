package frogstair

import "testing"

// TestFactorial tests the factorial function
func TestFactorial(t *testing.T) {
	data := []int{
		1,
		10,
		5,
		8,
	}
	expected := []int64{
		1,
		3628800,
		120,
		40320,
	}

	for i, d := range data {
		if expected[i] != Factorial(d) {
			t.Errorf("ERROR: Want %d, got %d", expected[i], Factorial(d))
		}
	}

}
