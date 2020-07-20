package frogstair

// Reverse reverses a string
func Reverse(s string) string {
	n := len(s)
	rs := make([]rune, n)
	for _, r := range s {
		n--
		rs[n] = r
	}
	return string(rs[n:])
}
