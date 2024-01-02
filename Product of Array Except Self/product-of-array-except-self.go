func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    res[0] = 1
    prefix := 1
    postfix := 1

    for i:=0; i<n-1; i++ {
        prefix *= nums[i]
        res[i+1] = prefix
    }
    for i:=n-1; i>0; i-- {
        postfix *= nums[i]
        res[i-1] *= postfix
    }

    return res
}