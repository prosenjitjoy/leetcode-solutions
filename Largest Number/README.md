# [Problem statement](https://leetcode.com/problems/largest-number)

Given a list of non-negative integers `nums`, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.

**Example 1:**


**Input:** nums = [10,2]
**Output:** "210"

**Example 2:**


**Input:** nums = [3,30,34,5,9]
**Output:** "9534330"

**Constraints:**

* `1 <= nums.length <= 100`
* `0 <= nums[i] <= 109`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1146025509/)

```go
func largestNumber(nums []int) string {
    numStr := make([]string, len(nums))
    for i, v := range nums {
        numStr[i] = strconv.Itoa(v)
    }

    sort.Slice(numStr, func(a, b int) bool {
        return numStr[a] + numStr[b] > numStr[b] + numStr[a]
    })

    ansStr := strings.Join(numStr, "")

    if ansStr[0] == '0' {
        return "0"
    }
    return ansStr
}
```