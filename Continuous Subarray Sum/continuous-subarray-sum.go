func checkSubarraySum(nums []int, k int) bool {
    m := map[int]int{0:-1}
    cumSum := 0

    for i, num := range nums {
        cumSum += num
        rem := cumSum%k
        if _, ok := m[rem]; !ok {
            m[rem] = i
        } else if i - m[rem] > 1 {
            return true
        }
    }

    return false
}