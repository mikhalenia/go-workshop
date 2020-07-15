package mglinnik

// reverse ...
func Reverse(x string) string {
	runes := []rune(x)
	for i, z := 0, len(runes) - 1; i < z; i, z = i + 1, z - 1 {
runes[i], runes[z] = runes[z], runes[i]
	}
	return string(runes)
}
