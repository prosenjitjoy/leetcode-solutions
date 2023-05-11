func majorityElement(nums []int) int {
    m := map[int]int{}
    for _, num := range nums {
        m[num]++
    }
    max := -1
    major := -1
    for key, val := range m {
        if val > max {
            max = val
            major = key
        }
    }
    return major
}