func subarraySum(nums []int, k int) int {
    prefixSum := map[int]int{0: 1}
    sum := 0
    res := 0
    
    for i:=0; i<len(nums); i++ {
        sum += nums[i]
        if v, ok := prefixSum[sum - k]; ok {
            res += v
        }
        prefixSum[sum]++
    }

    return res
}