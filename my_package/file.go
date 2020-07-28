package my_package

// Anagram
// str="I love the Beatles", secret="I love you"
func IsAnagram(line1 string, secret string) bool {
	x := make(map[int32]int)
	for _, letter := range line1 {
		if _, ok := x[letter]; ok {
			x[letter]++
		} else {
			x[letter] = 1
		}
	}

	for _, letter := range secret {
		if x[letter] < 0 {
			return false
		}

	}
	return true
}
