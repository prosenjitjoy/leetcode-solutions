# [Problem statement](https://leetcode.com/problems/minimum-penalty-for-a-shop)

You are given the customer visit log of a shop represented by a **0-indexed** string `customers` consisting only of characters `'N'` and `'Y'`:

* if the `ith` character is `'Y'`, it means that customers come at the `ith` hour
* whereas `'N'` indicates that no customers come at the `ith` hour.

If the shop closes at the `jth` hour (`0 <= j <= n`), the **penalty** is calculated as follows:

* For every hour when the shop is open and no customers come, the penalty increases by `1`.
* For every hour when the shop is closed and customers come, the penalty increases by `1`.

Return _the **earliest** hour at which the shop must be closed to incur a **minimum** penalty._

**Note** that if a shop closes at the `jth` hour, it means the shop is closed at the hour `j`.

**Example 1:**


**Input:** customers = "YYNY"
**Output:** 2
**Explanation:** 
- Closing the shop at the 0th hour incurs in 1+1+0+1 = 3 penalty.
- Closing the shop at the 1st hour incurs in 0+1+0+1 = 2 penalty.
- Closing the shop at the 2nd hour incurs in 0+0+0+1 = 1 penalty.
- Closing the shop at the 3rd hour incurs in 0+0+1+1 = 2 penalty.
- Closing the shop at the 4th hour incurs in 0+0+1+0 = 1 penalty.
Closing the shop at 2nd or 4th hour gives a minimum penalty. Since 2 is earlier, the optimal closing time is 2.

**Example 2:**


**Input:** customers = "NNNNN"
**Output:** 0
**Explanation:** It is best to close the shop at the 0th hour as no customers arrive.

**Example 3:**


**Input:** customers = "YYYY"
**Output:** 4
**Explanation:** It is best to close the shop at the 4th hour as customers arrive at each hour.

**Constraints:**

* `1 <= customers.length <= 105`
* `customers` consists only of characters `'Y'` and `'N'`.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1186472996/)

```go
func bestClosingTime(customers string) int {
    n := len(customers)+1
    prefix := 0
    postfix := 0
    prefix_no := make([]int, n)
    postfix_yes := make([]int, n)

    for i, j := 1, n-2; i<n; i, j = i+1, j-1 {
        prefix_no[i] = prefix
        if customers[i-1] == 'N' {
            prefix++
        }

        if customers[j] == 'Y' {
            postfix++
        }
        postfix_yes[j] = postfix

    }

    min := math.MaxInt
    j := 0
    for i:=0; i<n; i++ {
        penalty := prefix_no[i] + postfix_yes[i]
        if penalty < min {
            min = penalty
            j = i
        }
    }

    return j
}
```