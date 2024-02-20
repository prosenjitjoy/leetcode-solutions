func arraySign(nums []int) int {
    sign := 1

    for i := range nums {
        if nums[i] < 0 {
            sign *= -1
        } else if nums[i] == 0 {
            return 0
        }
    }

    return sign
}