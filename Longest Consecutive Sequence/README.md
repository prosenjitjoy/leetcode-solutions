# [Problem statement](https://leetcode.com/problems/longest-consecutive-sequence)

Given an unsorted array of integers `nums`, return _the length of the longest consecutive elements sequence._

You must write an algorithm that runs in `O(n)` time.

**Example 1:**


**Input:** nums = [100,4,200,1,3,2]
**Output:** 4
**Explanation:** The longest consecutive elements sequence is `[1, 2, 3, 4]`. Therefore its length is 4.

**Example 2:**


**Input:** nums = [0,3,7,2,5,8,4,6,0,1]
**Output:** 9

**Constraints:**

* `0 <= nums.length <= 105`
* `-109 <= nums[i] <= 109`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1137767234/)

```go
func longestConsecutive(nums []int) int {
    set := map[int]bool{}
    for _, num := range nums {
        set[num] = true
    }
    
    max := 0
    freq := map[int]int{}
    for k := range set {
        if !set[k-1] {
            i := k
            for set[i] {
                freq[k]++
                i++
            }

            if freq[k] > max {
                max = freq[k]
            }
        }
    }

    return max
}
```