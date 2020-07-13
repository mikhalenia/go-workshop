package string_reverse

import "testing"

func testFactorialResult(x, expected string, t *testing.T) {
	fibResult := Reverse(x)
	if fibResult != expected {
		t.Errorf("Unexpected result: %v, we expect: %v", fibResult, expected)
	}
}

func TestReverse(t *testing.T) {
	testFactorialResult("привет", "тевирп", t)
	testFactorialResult("hello", "olleh", t)
	testFactorialResult("programming go", "og gnimmargorp", t)
	testFactorialResult("1&$%^^   &*&^^&*", "*&^^&*&   ^^%$&1", t)
}

