package Test

// Fib(x) = Fib(x - 1) + Fib(x - 2)
// F(0) = 0
// F(1) = 1
func Fib(x int) int {
	if x < 2 {
		return x
	}

	return Fib(x - 1) + Fib(x - 2)
}
