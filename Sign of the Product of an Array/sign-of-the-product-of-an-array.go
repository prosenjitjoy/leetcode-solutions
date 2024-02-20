func arraySign(nums []int) int {
    negativeCount := 0
    n := len(nums)

    for i:=0; i<n; i++ {
        if nums[i] == 0 {
            return 0
        } else if nums[i] < 0 {
            negativeCount++
        }
    }

    if negativeCount%2 == 0 {
        return 1
    } else {
        return -1
    }
}