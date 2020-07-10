package mglinnik

import "testing"

const failure = "Expected %d to equal %d, but got %d"

// TestFactorial ...
func TestFactorial(t *testing.T) {
	input := []int{1, 5, 7, 10}
	expected := []int{1, 120, 5040, 3628800}

	for i, x := range input {
		if Factorial(x) != expected[i] {
			t.Errorf(failure, input[i], expected[i], Factorial(x))

		}

	}

}
