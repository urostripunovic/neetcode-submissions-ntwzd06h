func removeElement(nums []int, val int) int {
    pointer := 0
	for _, num := range nums {
		if num != val {
			nums[pointer] = num
			pointer++
		}
	}
	return pointer
}
