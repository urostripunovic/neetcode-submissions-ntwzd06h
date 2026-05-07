func isValid(s string) bool {
	correct := false
	stack := Stack[rune]{}
	closingChars := []rune{')', ']', '}'}
	for _, ch := range s {
		//insert
		fmt.Println(ch)
		//before inserting and item is closing bracket
		if !Some(closingChars, func(closingChar rune) bool {
			return ch == closingChar
		}) {
			stack.Push(ch)
			continue
		}
		//peek to see if it is opening
		top, _ := stack.Peek()
		//if true pop, else continue
		if top + 2 == ch {
			stack.Pop()
		} else if top + 1 == ch {
			stack.Pop()
		} else {
			stack.Push(ch)
		}
	}
	if len(stack.items) == 0 {
		correct = !correct
	}
	return correct
}

func Some[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

type Stack [T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if (len(s.items) == 0) {
		var zero T
		return zero, false
	}
	topIndex := len(s.items) - 1
	val := s.items[topIndex]
	s.items = s.items[:topIndex]
	return val, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if (len(s.items) == 0) {
		var zero T
		return zero, false
	}

	topIndex := len(s.items) - 1
	val := s.items[topIndex]
	return val, true
}
