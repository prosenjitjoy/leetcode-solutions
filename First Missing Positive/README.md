# [Problem statement](https://leetcode.com/problems/first-missing-positive)

Given an unsorted integer array `nums`, return the smallest missing positive integer.

You must implement an algorithm that runs in `O(n)` time and uses `O(1)` auxiliary space.

**Example 1:**


**Input:** nums = [1,2,0]
**Output:** 3
**Explanation:** The numbers in the range [1,2] are all in the array.

**Example 2:**


**Input:** nums = [3,4,-1,1]
**Output:** 2
**Explanation:** 1 is in the array but 2 is missing.

**Example 3:**


**Input:** nums = [7,8,9,11,12]
**Output:** 1
**Explanation:** The smallest positive integer 1 is missing.

**Constraints:**

* `1 <= nums.length <= 105`
* `-231 <= nums[i] <= 231 - 1`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1179552864/)

```go
func firstMissingPositive(nums []int) int {
    set := map[int]bool{}
    n := len(nums)

    for i:=0; i<n; i++ {
        set[nums[i]] = true
    }

    n++
    for i:=1; i<=n; i++ {
        if _, ok := set[i]; !ok {
            return i
        }
    }

    return n
}
```