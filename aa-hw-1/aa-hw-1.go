package aahw1

import (
	"errors"
	"strconv"
)

const FOOBAR, FOO, BAR = "foobar", "foo", "bar"

// ### FACTORIAL ####################################################################

func Factorial1(val int) (result int, err error) {
	result = 1

	switch true {
	case val < 0:
		err = errors.New("Value less than 0")
	default:
		for i := 1; i <= val; i++ {
			result = result * i
		}
		if result < 0 {
			err = errors.New("The result value more than Integer type can contain")
		}
	}
	return
}

func Factorial2(val int) (result int) {

	switch true {
	case val < 0 || val > 20:
		result = -1
	case val >= 0 && val <= 1:
		result = 1
	default:
		result = val * Factorial2(val-1)
	}
	return
}

// ### FOOBAR ####################################################################

func FoobarV1(val int) (result string) {
	result = ""

	if val%3 == 0 && val%5 == 0 {
		result = strconv.Itoa(val)
	} else {
		if val%3 == 0 {
			result = FOO
		}

		if val%5 == 0 {
			result = BAR
		}
	}
	return
}

func FoobarV2(val int) (result string) {
	switch true {
	case val%3 == 0 && val%5 == 0:
		result = strconv.Itoa(val)
	case val%3 == 0:
		result = FOO
	case val%5 == 0:
		result = BAR
	default:
		result = ""
	}
	return
}

// ### REVERS ####################################################################

func Revers(s string) string {
	sToRune = []rune(s)
	ln := len(sToRune)
	result := make([]rune, ln)
	for i, val := range sToRune {
		result[ln-i-1] = val
	}
	return string(result)
}
