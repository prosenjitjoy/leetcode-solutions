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

# [Solution in go](https://leetcode.com/submissions/detail/948415767/)

```go
func findDisappearedNumbers(nums []int) []int {
    ans := []int{}
    for _, num := range nums {
        nums[abs(num)-1] = abs(nums[abs(num)-1]) * -1
    }
    for idx, num := range nums {
        if num > 0 {
            ans = append(ans, idx+1)
        }
    }
    return ans
}

func abs(num int) int {
    if num < 0 { 
        return num * -1
    } else {
        return num
    }
}
```