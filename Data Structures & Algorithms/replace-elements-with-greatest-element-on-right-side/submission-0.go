func replaceElements(arr []int) []int {
    maxSoFar := -1
    for index := len(arr) -1 ; index >= 0; index-- {
        tmp := arr[index]
        arr[index] = maxSoFar
        if tmp > maxSoFar {
            maxSoFar = tmp
        }
    }

    return arr
}
