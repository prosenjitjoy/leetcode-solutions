# [Problem statement](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array)

Given an array `nums` of `n` integers where `nums[i]` is in the range `[1, n]`, return _an array of all the integers in the range_ `[1, n]` _that do not appear in_ `nums`.

**Example 1:**

**Input:** nums = [4,3,2,7,8,2,3,1]
**Output:** [5,6]

**Example 2:**

**Input:** nums = [1,1]
**Output:** [2]

**Constraints:**

* `n == nums.length`
* `1 <= n <= 105`
* `1 <= nums[i] <= n`

**Follow up:** Could you do it without extra space and in `O(n)` runtime? You may assume the returned list does not count as extra space.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/948419148/)

```go
func findDisappearedNumbers(nums []int) []int {
    ans := []int{}
    found := make([]bool, len(nums))

    for _, num := range nums {
        found[num-1] = true
    }
    for i, v := range found {
        if v == false {
            ans = append(ans, i+1)
        }
    }
    return ans
}
```