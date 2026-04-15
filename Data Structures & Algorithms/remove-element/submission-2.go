func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			right--
			nums[left] = nums[right]
		} else {
			left++
		}
	}
	return left
}
