package dudkin

/*
#1 Напишите функцию проверки балансировки скобок
пример:
строка: ({{abc[aa],}-})
вывод: true
*/


/*
#2 Напишите функцию поиска минимального элемента в бинарном дереве
 */

func isOpen(c rune) bool {
	return c == '{' || c == '[' || c == '('
}

func isClosed(c rune) bool {
	return c == '}' || c == ']' || c == ')'
}

type stack struct {
	data []rune
}

func (s *stack)Push(b rune) {
	s.data = append(s.data, b)
}

func (s *stack)Pop() rune {
	b := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return b
}

func (s *stack)Length() int {
	return len(s.data)
}

func Balanced(data string) bool {
	s := new(stack)

	for _, ch := range data {
		if isOpen(ch) {
			s.Push(ch)
		} else if isClosed(ch) {
			s.Pop()
		}
	}

	return s.Length() == 0
}