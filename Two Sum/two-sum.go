func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, num := range nums {
        val, ok := m[target - num]
        if ok {
            return []int{i, val}
        } else {
            m[num] = i
        }
    }
    return []int{}
}