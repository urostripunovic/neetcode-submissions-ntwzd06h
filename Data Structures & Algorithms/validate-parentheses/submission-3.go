func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}

	stack := Stack[rune]{}
	closers := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

    for _, ch := range s {
		if opener, ok := closers[ch]; ok {
			if stack.Empty() {
				return false 
			} 
			top, notEmpty := stack.Pop()
			if notEmpty && top != opener {
				return false
			}
		} else {
			stack.Push(ch)
		}
	}
	return stack.Empty()
}

type Stack[T any] struct {
	items []T
}

func (this *Stack[T]) Push(val T) {
	this.items = append(this.items, val)
}

func (this *Stack[T]) Pop() (T, bool) {
	if len(this.items) == 0 {
		var zero T
		return zero, false
	}
	topIndex := len(this.items) - 1
	val := this.items[topIndex]
	this.items = this.items[:topIndex]
	return val, true
}

func (this *Stack[T]) Empty() bool {
	return len(this.items) == 0
}
