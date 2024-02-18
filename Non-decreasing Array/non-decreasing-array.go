func checkPossibility(nums []int) bool {
    changed := false
    n := len(nums)-1
    
    for i:=0; i<n; i++ {
        if nums[i] > nums[i+1] {
            if changed {
                return false
            }

            if i==0 || nums[i+1] >= nums[i-1] {
                nums[i] = nums[i+1]
            } else {
                nums[i+1] = nums[i]
            }

            changed = true
        }
    }

    return true
}