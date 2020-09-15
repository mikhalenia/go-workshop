package averianov

/*
#1 Напишите функцию поиска наименьшего числа в отсортированном массиве, которое делится на 3
 */

/*
#2 Напишите функцию выводящую количество заданных символов в строке
пример:
строка: AndOdddPoSd
символы: [A, d]
вывод:
A: 1
d: 5
 */

func iterateRuns(str string, runes []rune) map[rune]int {
	result := map[rune]int{}

	for _, r := range runes {
		for _, s := range str {
			if s == r {
				result[r]++
			}
		}
	}

	return result
}

func leastInArray(arr []int) (result int) {

	result = -1
	if arr[0] < arr[len(arr)-1] {
		i := 0
		for {
			if arr[i]%3 == 0 {
				val := arr[i] / 3
				if result > val || result == -1 {
					result = arr[i]
					return
				}
			} else {
				i++
				continue
			}
		}
	} else {
		i := len(arr) - 1
		for {
			if arr[i]%3 == 0 {
				val := arr[i] / 3
				if result > val || result == -1 {
					result = arr[i]
					return
				}
			} else {
				i--
				continue
			}
		}
	}

	return
}

