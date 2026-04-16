func replaceElements(arr []int) []int {
    maxSoFar := -1
    for index := len(arr) -1 ; index >= 0; index-- {
        arr[index], maxSoFar = maxSoFar, max(maxSoFar, arr[index])
    }

    return arr
}
