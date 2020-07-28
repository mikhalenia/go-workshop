package palindrome

// Palindrome returns if a string is a palindrome
func Palindrome(s string) bool {
	for i, c := range s {
		if uint8(c) != s[len(s) - i - 1] {
			return false
		}
	}
	return true
}