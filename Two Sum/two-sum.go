func twoSum(nums []int, target int) []int {
    m := make(map[int]int, 2)
    for i, num := range nums {
        if val, ok := m[target - num]; ok {
            return []int{i, val}
        }
        m[num] = i
    }
    return nil
}