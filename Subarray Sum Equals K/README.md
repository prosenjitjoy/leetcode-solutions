# [Problem statement](https://leetcode.com/problems/subarray-sum-equals-k)

Given an array of integers `nums` and an integer `k`, return _the total number of subarrays whose sum equals to_ `k`.

A subarray is a contiguous **non-empty** sequence of elements within an array.

**Example 1:**

**Input:** nums = [1,1,1], k = 2
**Output:** 2

**Example 2:**

**Input:** nums = [1,2,3], k = 3
**Output:** 2

**Constraints:**

* `1 <= nums.length <= 2 * 104`
* `-1000 <= nums[i] <= 1000`
* `-107 <= k <= 107`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1141815884/)

```go
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
```