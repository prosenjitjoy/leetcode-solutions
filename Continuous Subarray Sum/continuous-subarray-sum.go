func checkSubarraySum(nums []int, k int) bool {
    m := map[int]int{0:-1}
    cumSum := 0

    for i, num := range nums {
        cumSum += num
        rem := cumSum%k
        if v, ok := m[rem]; ok {
            if i-v > 1 {
                return true
            }
        } else {
            m[rem] = i
        }
    }

    return false
}