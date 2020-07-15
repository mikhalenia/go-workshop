package factorial

func Factorial(number int)(result int) {
	result = 1
	if number > 0 {
		result = number * Factorial(number-1)
		return result
	}
	return result
}
