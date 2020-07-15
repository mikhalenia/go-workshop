package mglinnik

import "strconv"

// Foobar ...
func Foobar(x int) string {
	if x%3 == 0 && x%5 == 0 {
		j := strconv.Itoa(x)
		return j
	}
	if x%3 == 0 {
		return "foo"
	}
	if x%5 == 0 {
		return "bar"
	}
	return ""
}
