# [Problem statement](https://leetcode.com/problems/find-the-difference-of-two-arrays)

Given two **0-indexed** integer arrays `nums1` and `nums2`, return _a list_ `answer` _of size_ `2` _where:_

* `answer[0]` _is a list of all **distinct** integers in_ `nums1` _which are **not** present in_ `nums2`_._
* `answer[1]` _is a list of all **distinct** integers in_ `nums2` _which are **not** present in_ `nums1`.

**Note** that the integers in the lists may be returned in **any** order.

**Example 1:**


**Input:** nums1 = [1,2,3], nums2 = [2,4,6]
**Output:** [[1,3],[4,6]]
**Explanation:**
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].

**Example 2:**


**Input:** nums1 = [1,2,3,3], nums2 = [1,1,2,2]
**Output:** [[3],[]]
**Explanation:**
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].

**Constraints:**

* `1 <= nums1.length, nums2.length <= 1000`
* `-1000 <= nums1[i], nums2[i] <= 1000`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1181789160/)

```go
func findDifference(nums1 []int, nums2 []int) [][]int {
    numSet1 := map[int]bool{}
    numSet2 := map[int]bool{}

    for i := range nums1 {
        numSet1[nums1[i]] = true
    }
    for i := range nums2 {
        numSet2[nums2[i]] = true
    }

    ans1 := []int{}
    ans2 := []int{}

    for key := range numSet1 {
        if !numSet2[key] {
            ans1 = append(ans1, key)
        }
    }
    for key := range numSet2 {
        if !numSet1[key] {
            ans2 = append(ans2, key)
        }
    }

    return [][]int{ans1, ans2}
}
```