# [Problem statement](https://leetcode.com/problems/contains-duplicate)

Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.

**Example 1:**

**Input:** nums = [1,2,3,1]
**Output:** true

**Example 2:**

**Input:** nums = [1,2,3,4]
**Output:** false

**Example 3:**

**Input:** nums = [1,1,1,3,3,4,3,2,4,2]
**Output:** true

**Constraints:**

* `1 <= nums.length <= 105`
* `-109 <= nums[i] <= 109`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/947721358/)

```go
func containsDuplicate(nums []int) bool {
    m := map[int]int{}
    for _, num := range nums {
        if _, ok := m[num]; ok {
            return true
        }
        m[num]++
    }
    return false
}
```