func sortColors(nums []int)  {
    left := 0
    right := len(nums) - 1
    
    for i := 0; i<=right; i++ {
        if nums[i] == 0 {
            nums[i], nums[left] = nums[left], nums[i]
            left++
        }
        if nums[i] == 2 {
            nums[i], nums[right] = nums[right], nums[i]
            right--
            i--
        }
    }
}