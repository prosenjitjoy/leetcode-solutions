# [Problem statement](https://leetcode.com/problems/continuous-subarray-sum)

Given an integer array nums and an integer k, return `true` _if_ `nums` _has a **good subarray** or_ `false` _otherwise_.

A **good subarray** is a subarray where:

* its length is **at least two**, and
* the sum of the elements of the subarray is a multiple of `k`.

**Note** that:

* A **subarray** is a contiguous part of the array.
* An integer `x` is a multiple of `k` if there exists an integer `n` such that `x = n * k`. `0` is **always** a multiple of `k`.

**Example 1:**


**Input:** nums = [23,2,4,6,7], k = 6
**Output:** true
**Explanation:** [2, 4] is a continuous subarray of size 2 whose elements sum up to 6.

**Example 2:**


**Input:** nums = [23,2,6,4,7], k = 6
**Output:** true
**Explanation:** [23, 2, 6, 4, 7] is an continuous subarray of size 5 whose elements sum up to 42.
42 is a multiple of 6 because 42 = 7 * 6 and 7 is an integer.

**Example 3:**


**Input:** nums = [23,2,6,4,7], k = 13
**Output:** false

**Constraints:**

* `1 <= nums.length <= 105`
* `0 <= nums[i] <= 109`
* `0 <= sum(nums[i]) <= 231 - 1`
* `1 <= k <= 231 - 1`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1147272685/)

```go
func checkSubarraySum(nums []int, k int) bool {
    m := map[int]int{0:-1}
    cumSum := 0

    for i, num := range nums {
        cumSum += num
        rem := cumSum%k
        if _, ok := m[rem]; !ok {
            m[rem] = i
        } else if i - m[rem] > 1 {
            return true
        }
    }

    return false
}
```