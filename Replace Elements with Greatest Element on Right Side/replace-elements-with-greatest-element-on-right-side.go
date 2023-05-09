func replaceElements(arr []int) []int {
    length := len(arr)
    ans := make([]int, length)
    max := arr[length-1]

    for i := length-1; i>=0; i-- {
        if arr[i] > max {
            max = arr[i]
        }
        ans[i] = max
    }
    ans = ans[1:]
    ans = append(ans, -1)
    return ans
}