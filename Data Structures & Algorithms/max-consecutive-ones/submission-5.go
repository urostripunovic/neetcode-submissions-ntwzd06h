func findMaxConsecutiveOnes(nums []int) int {
    curr_longest, longest := 0, 0
    for _, val := range nums {
        if val == 1 {
            curr_longest++
        } else {
            curr_longest = 0
        }

        if curr_longest > longest {
            longest = curr_longest
        }
    }
    return longest
}
