func replaceElements(arr []int) []int {
    max := -1
    for i := len(arr)-1; i >= 0; i-- {
        tmp := max
        if arr[i] > max {
            max = arr[i]
        }
        arr[i] = tmp
    }
    return arr
}
