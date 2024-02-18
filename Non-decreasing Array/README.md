# [Problem statement](https://leetcode.com/problems/non-decreasing-array)

Given an array `nums` with `n` integers, your task is to check if it could become non-decreasing by modifying **at most one element**.

We define an array is non-decreasing if `nums[i] <= nums[i + 1]` holds for every `i` (**0-based**) such that (`0 <= i <= n - 2`).

**Example 1:**


**Input:** nums = [4,2,3]
**Output:** true
**Explanation:** You could modify the first 4 to 1 to get a non-decreasing array.

**Example 2:**


**Input:** nums = [4,2,1]
**Output:** false
**Explanation:** You cannot get a non-decreasing array by modifying at most one element.

**Constraints:**

* `n == nums.length`
* `1 <= n <= 104`
* `-105 <= nums[i] <= 105`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1178980910/)

```go
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
```