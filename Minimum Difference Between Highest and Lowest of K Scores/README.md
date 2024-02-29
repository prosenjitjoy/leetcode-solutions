# [Problem statement](https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores)

You are given a **0-indexed** integer array `nums`, where `nums[i]` represents the score of the `ith` student. You are also given an integer `k`.

Pick the scores of any `k` students from the array so that the **difference** between the **highest** and the **lowest** of the `k` scores is **minimized**.

Return _the **minimum** possible difference_.

**Example 1:**


**Input:** nums = [90], k = 1
**Output:** 0
**Explanation:** There is one way to pick score(s) of one student:
- [**90**]. The difference between the highest and lowest score is 90 - 90 = 0.
The minimum possible difference is 0.

**Example 2:**


**Input:** nums = [9,4,1,7], k = 2
**Output:** 2
**Explanation:** There are six ways to pick score(s) of two students:
- [**9**,**4**,1,7]. The difference between the highest and lowest score is 9 - 4 = 5.
- [**9**,4,**1**,7]. The difference between the highest and lowest score is 9 - 1 = 8.
- [**9**,4,1,**7**]. The difference between the highest and lowest score is 9 - 7 = 2.
- [9,**4**,**1**,7]. The difference between the highest and lowest score is 4 - 1 = 3.
- [9,**4**,1,**7**]. The difference between the highest and lowest score is 7 - 4 = 3.
- [9,4,**1**,**7**]. The difference between the highest and lowest score is 7 - 1 = 6.
The minimum possible difference is 2.

**Constraints:**

* `1 <= k <= nums.length <= 1000`
* `0 <= nums[i] <= 105`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1189477176/)

```go
func minimumDifference(nums []int, k int) int {
    nums = mergeSort(nums)
    res := math.MaxInt
    
    for l,r := 0,k-1; r < len(nums); l,r = l+1,r+1 {
        res = min(res, nums[r] - nums[l])
    }

    return res
}

func mergeSort(items []int) []int {
    if len(items) < 2 {
        return items
    }

    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])

    return merge(first, second)
}

func merge(a, b []int) []int {
    final := []int{}
    i,j := 0, 0

    for i<len(a) && j<len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }

    for i<len(a) {
        final = append(final, a[i])
        i++
    }
    for j<len(b) {
        final = append(final, b[j])
        j++
    }

    return final
}
```