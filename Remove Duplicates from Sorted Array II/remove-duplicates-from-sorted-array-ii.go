func removeDuplicates(nums []int) int {
    n := len(nums)
    l, r := 0, 0
    for r < n {
        count := 1
        for r+1 < n && nums[r] == nums[r+1] {
            r++
            count++
        }
        m := min(2, count)
        for i:=0; i<m; i++ {
            nums[l] = nums[r]
            l++
        }
        r++
    }

    return l
}