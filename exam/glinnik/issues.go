package glinnik

/*
#1 Напишите функцию проверки полиндрома
пример:
строка: abba
вывод: true
*/
func Palindrome(data string) bool {
	for i := 0; i < len(data) / 2; i++ {
		if data[i] != data[len(data)-i-1] {
			return false
		}
	}
	return true
}

/*
#2 Напишите функцию поиска минимального элемента в графе
*/

/*
1 - 3
 \   \
 2 - 5 - 4
{
	1: [...],
	2: [...],
	3: [...],
	4: [...],
	5: [...]
}
 */

func MinGraphElement(G map[int]int) int {
/*	(*visited)[node] = true
for _, el := range (*G) {
	if _, ok := (*visited)[el]; !ok {
		visitedNode := MinGraphElement(G, el, visited)

		}
	}
}

 */
var minNumber int
for min := range G {
	minNumber = min
	break
}
for min := range G {
	if min < minNumber {
		minNumber = min
	}
}
return minNumber
}



