func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, num := range nums {
        if val, ok := m[target - num]; ok {
            return []int{i, val}
        }
        m[num] = i
    }
    return nil
}