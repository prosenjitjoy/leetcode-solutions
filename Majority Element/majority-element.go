func majorityElement(nums []int) int {
    count := -1
    res := -1

    for _, num := range nums {
        if count < 0 {
            res = num
            count++
        } else if num == res {
            count++
        } else {
            count--
        }
    }
    return res
}