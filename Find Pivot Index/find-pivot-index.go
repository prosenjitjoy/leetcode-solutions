func pivotIndex(nums []int) int {
    ans := -1
    cum := make([]int, len(nums))
    cum[0] = nums[0]

    for i:=1; i<len(nums); i++ {
        cum[i] = cum[i-1]+nums[i]
    }
    for i, num := range nums {
        left := cum[i] - num
        right := cum[len(cum)-1] - cum[i]
        if left == right {
            ans = i
            break
        }
    }
    return ans
}