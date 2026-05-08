func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}

	stack := Stack[rune]{
		items: make([]rune, 0, len(s)),
	}

    for _, ch := range s {
		switch ch {
			case ')':
				if stack.Empty() { return false } 
				if top, ok := stack.Peek(); ok && top != '(' {
					return false
				}
				stack.Pop()
			case ']':
				if stack.Empty() { return false } 
				if top, ok := stack.Peek(); ok && top != '[' {
					return false
				}
				stack.Pop()
			case '}':
				if stack.Empty() { return false } 
				if top, ok := stack.Peek(); ok && top != '{' {
					return false
				}
				stack.Pop()
			default:
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

func (this *Stack[T]) Peek() (T, bool) {
	if len(this.items) == 0 {
		var zero T
		return zero, false
	}
	topIndex := len(this.items) - 1
	val := this.items[topIndex]
	return val, true
}

func (this *Stack[T]) Empty() bool {
	return len(this.items) == 0
}
