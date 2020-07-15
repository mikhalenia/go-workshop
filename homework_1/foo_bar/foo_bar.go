package foo_bar

import "strconv"

func FooBar(number int)(result string) {
	if number % 3 == 0 {
		result = "foo"
	}
	if number % 5 == 0 {
		result = "bar"
	}
	if number % 3 == 0 && number % 5 == 0 {
		result = strconv.Itoa(number)
	}
	return result
}