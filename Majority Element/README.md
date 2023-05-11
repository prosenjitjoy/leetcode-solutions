# [Problem statement](https://leetcode.com/problems/majority-element)

Given an array `nums` of size `n`, return _the majority element_.

The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

**Example 1:**

**Input:** nums = [3,2,3]
**Output:** 3

**Example 2:**

**Input:** nums = [2,2,1,1,1,2,2]
**Output:** 2

**Constraints:**

* `n == nums.length`
* `1 <= n <= 5 * 104`
* `-109 <= nums[i] <= 109`

**Follow-up:** Could you solve the problem in linear time and in `O(1)` space?

<br />

# [Solution in go](https://leetcode.com/submissions/detail/948182566/)

```go
func majorityElement(nums []int) int {
    m := map[int]int{}
    for _, num := range nums {
        m[num]++
    }
    max := -1
    major := -1
    for key, val := range m {
        if val > max {
            max = val
            major = key
        }
    }
    return major
}
```