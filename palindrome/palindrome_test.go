package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	data := []string{
		"radar",
		"rotator",
		"racecar",
		"ne palindrom",
	}

	expected := []bool {
		true,
		true,
		true,
		false,
	}

	for i, s := range data {
		if expected[i] != Palindrome(s) {
			t.Errorf("ERROR: Expected: %v, got: %v", expected[i], Palindrome(s))
		}
	}
}