func calPoints(operations []string) int {
	s := []int{}
	size := 0
	for _, o := range operations {
		if o == "C" {
			size--
			s = s[:size]
		} else if o == "D" {
			s = append(s, s[size - 1] * 2)
			size++
		} else if o == "+" {
			s = append(s, s[size - 2] + s[size - 1])
			size++
		} else {
			i, _ := strconv.Atoi(o)
			s = append(s, i)
			size++
		}
	}

	sum := 0
	for _, val := range s {
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