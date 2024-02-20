func arraySign(nums []int) int {
    sign := 1
    n := len(nums)

    for i:=0; i<n; i++ {
        if nums[i] < 0 {
            sign *= -1
        } else if nums[i] == 0 {
            return 0
        }
    }

    return sign
}