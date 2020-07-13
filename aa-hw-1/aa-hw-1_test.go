package aahw1

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	wantValue := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800,
		479001600, 6227020800, 87178291200, 1307674368000, 20922789888000, 355687428096000,
		6402373705728000, 121645100408832000, 2432902008176640000, 21}

	// Test function Factorial1
	for i := -1; i <= 21; i++ {
		if result, err := Factorial1(i); err != nil {
			if i < 20 && i >= 0 {
				t.Errorf("Error Test Factorial1. Iteration = %d, Want = %d and Got = '%v'.", i, wantValue[i], err)
			} else if err == nil {
				t.Errorf("Error Test Factorial1. Error must be not nil")
			}
		} else if result != wantValue[i] {
			t.Errorf("Error Test Factorial1. Iteration = %d, Want = %d and Got = '%v'.", i, wantValue[i], result)
		}
	}

	// Test function Factorial2
	for i := -1; i <= 21; i++ {
		result := Factorial2(i)
		if i >= 0 && i <= 20 { // if i >= 21 than Integer can`t contain result
			if result != wantValue[i] {
				t.Errorf("Error Test Factorial2. Iteration = %d, Want = %d and Got = '%v'.", i, wantValue[i], result)
			}
		} else {
			if result != -1 {
				t.Errorf("Error Test Factorial2. Result must be -1. Iteration = %d, Want = %d and Got = '%v'", i, wantValue[i], result)
			}
		}
	}
}

func TestFoobar(t *testing.T) {

	setValue := []int{12, 35, 75, 678, 18, 8}
	wantValue := []string{FOO, BAR, "75", FOO, FOO, ""}
	var result string

	// Test function Foobar1
	for i, val := range setValue {
		result = FoobarV1(val)

		if result != wantValue[i] {
			t.Errorf("Error Test FoobarV1. Want = %s and Got = %s.", wantValue[i], result)
		}
	}

	// Test function Foobar2
	for i, val := range setValue {
		result = FoobarV2(val)

		if result != wantValue[i] {
			t.Errorf("Error Test FoobarV2. Want = %s and Got = %s.", wantValue[i], result)
		}
	}

}

func TestRevers(t *testing.T) {

	setValue := []string{"tceles", "select", "57", "зазор", "кул", ""}
	wantValue := []string{"select", "tceles", "75", "розаз", "лук", ""}
	var result string

	// Test function Revers
	for i, val := range setValue {
		result = Revers(val)

		if result != wantValue[i] {
			t.Errorf("Error Test Revers. Want = %s and Got = %s.", wantValue[i], result)
		}
	}
}
