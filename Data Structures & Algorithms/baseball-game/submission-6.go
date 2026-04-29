func calPoints(operations []string) int {
	stack := make([]int, len(operations)) 
    size := 0 // This acts as our "Stack Pointer"
	sum := 0

	for _, operation := range operations {
		if operation == "C" {
			size--
			sum -= stack[size]
		} else if operation == "D" {
			stack[size] = stack[size - 1] * 2
			size++
			sum += stack[size - 1]
		} else if operation == "+" {
			stack[size] = stack[size - 2] + stack[size - 1]
			size++
			sum += stack[size - 1]
		} else {
			num, _ := strconv.Atoi(operation)
			stack[size] = num
			size++
			sum += stack[size - 1]
		}
	}
	return sum
}
