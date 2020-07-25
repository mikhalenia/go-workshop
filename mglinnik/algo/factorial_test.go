package mglinnik

import "testing"

const failureInteger = "Expected %d to equal %d, but got %d"

// TestFactorial ...
func TestFactorial(t *testing.T) {
	input := []int{1, 5, 7, 10}
	expected := []int{1, 120, 5040, 3628800}

	for i, x := range input {
		if Factorial(x) != expected[i] {
			t.Errorf(failureInteger, input[i], expected[i], Factorial(x))

		}

	}

}
