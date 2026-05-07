func isValid(s string) bool {
	stack := Stack[rune]{}
	closeToOpen := map[rune]rune{
    	')': '(',
    	']': '[',
    	'}': '{',
	}

	for _, ch := range s {
		if opener, exists := closeToOpen[ch]; exists {
			if !stack.Empty() {
				top, ok := stack.Pop()
				if ok && top != opener {
					return false
				}
			} else {
				return false
			}
		} else {
			stack.Push(ch)
		}
		
	}
	return stack.Empty()
}

type Stack [T any] struct {
	items []T
}

func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
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
