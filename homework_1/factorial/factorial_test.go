package factorial

import "testing"

func testFactorialResult(x, expected int, t *testing.T) {
	fibResult := Factorial(x)
	if fibResult != expected {
		t.Errorf("Unexpected result: %v, we expect: %v", fibResult, expected)
	}
}

func TestFactorial(t *testing.T) {
	testFactorialResult(0, 1, t)
	testFactorialResult(1, 1, t)
	testFactorialResult(2, 2, t)
	testFactorialResult(3, 6, t)
	testFactorialResult(4, 24, t)
	testFactorialResult(11, 39916800, t)
	testFactorialResult(20, 2432902008176640000, t)
}






