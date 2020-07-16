package aahw2

func hourglassSum(arr [][]int32) int32 {
	var max int32
	max = -9 * 7

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var sum int32
			sum = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func rotLeft(a []int32, d int32) []int32 {
	return append(a[d:], a[:d]...)
}

func alternatingCharacters(s string) int32 {
	var last rune
	result := int32(0)
	for _, v := range s {
		if v != last {
			last = v
		} else {
			result++
		}
	}
	return result
}
