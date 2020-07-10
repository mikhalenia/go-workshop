package frogstair

// Factorial returns a factorial of a number
func Factorial(x int) int64 {
	if x == 1 {
		return 1
	}
	return Factorial(x-1) * int64(x)
}
