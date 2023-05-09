func containsDuplicate(nums []int) bool {
    m := map[int]int{}
    for _, num := range nums {
        if _, ok := m[num]; ok {
            return true
        }
        m[num]++
    }
    return false
}