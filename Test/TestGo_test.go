package Test

import "testing"

func testFibResult(x, expected int, t *testing.T) {
	fibResult := Fib(x)
	if fibResult != expected {
		t.Errorf("Unexpected result: %v, we expect: %d", fibResult, expected)
	}
}

func TestFib(t *testing.T) {
	testFibResult(0, 0, t)
	testFibResult(1, 1, t)
	testFibResult(2, 1, t)
	testFibResult(3, 2, t)
	testFibResult(4, 3, t)
	testFibResult(5, 5, t)
	testFibResult(6, 8, t)
}
