func calPoints(operations []string) int {
	s := Stack[int]{}

	for _, o := range operations {
		if o == "C" {
			s.Pop()
			continue
		} else if o == "D" {
			v, _ := s.Peek()
			s.Push(v * 2)
			continue
		} else if o == "+" {
			left,_ := s.Peek()
			s.Pop()
			right,_ := s.Peek()
			s.Push(left)
			s.Push(left + right)
			continue
		}

		v := o
		i, _ := strconv.Atoi(v)
		s.Push(i)
	}

	sum := 0
	for _, val := range s.items {
		sum += val
	}

	return sum
}

type Stack [T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	top_index := len(s.items) - 1
	item := s.items[top_index]

	s.items = s.items[:top_index]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items) - 1], true
}