# [Problem statement](https://leetcode.com/problems/move-zeroes)

Given an integer array `nums`, move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.

**Note** that you must do this in-place without making a copy of the array.

**Example 1:**

**Input:** nums = [0,1,0,3,12]
**Output:** [1,3,12,0,0]

**Example 2:**

**Input:** nums = [0]
**Output:** [0]

**Constraints:**

* `1 <= nums.length <= 104`
* `-231 <= nums[i] <= 231 - 1`

**Follow up:** Could you minimize the total number of operations done?

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1192303600/)

```go
func moveZeroes(nums []int)  {
    l := 0
    for r:=0; r < len(nums); r++ {
        if nums[r] != 0 {
            nums[l], nums[r] = nums[r], nums[l]
            l++
        }
    }
}
```