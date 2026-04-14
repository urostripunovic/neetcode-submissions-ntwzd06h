func findMaxConsecutiveOnes(nums []int) int {
	//two pointer, one ahead compare to before
    curr_longest, longest := 0, 0
    for _, val := range nums {
        if val == 1 {
            curr_longest++
        } else {
            curr_longest = 0
        }
        longest = max(longest, curr_longest)
    }
    return longest
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
