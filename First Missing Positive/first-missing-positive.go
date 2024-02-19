func firstMissingPositive(nums []int) int {
    set := map[int]bool{}
    n := len(nums)

    for i:=0; i<n; i++ {
        set[nums[i]] = true
    }

    n++
    for i:=1; i<=n; i++ {
        if _, ok := set[i]; !ok {
            return i
        }
    }

    return n
}