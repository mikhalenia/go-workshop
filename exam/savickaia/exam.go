package savickaia

func MinNumber(mass []int) int {
	for _, x := range mass{
		if x%4 == 0 {
			return x
		}
	}
	return -1
}

func Anagram(line string, anagram string) bool {
	var dict = map[int32]int{}
	for _, ch := range line{
		_, ok := dict[ch]
		if ok {
			dict[ch]++
		} else {
			dict[ch] = 1
		}
	}

	for _, ch := range anagram{
		dict[ch]--
		if dict[ch] < 0{
			return false
		}
	}
	return true
}


