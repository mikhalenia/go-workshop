package string_reverse

import "testing"

func testReverseResult(x, expected string, t *testing.T) {
	fibResult := Reverse(x)
	if fibResult != expected {
		t.Errorf("Unexpected result: %v, we expect: %v", fibResult, expected)
	}
}

func TestReverse(t *testing.T) {
	testReverseResult("привет", "тевирп", t)
	testReverseResult("hello", "olleh", t)
	testReverseResult("programming go", "og gnimmargorp", t)
	testReverseResult("1&$%^^   &*&^^&*", "*&^^&*&   ^^%$&1", t)
}

