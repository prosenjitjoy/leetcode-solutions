# [Problem statement](https://leetcode.com/problems/sort-an-array)

Given an array of integers `nums`, sort the array in ascending order and return it.

You must solve the problem **without using any built-in** functions in `O(nlog(n))` time complexity and with the smallest space complexity possible.

**Example 1:**


**Input:** nums = [5,2,3,1]
**Output:** [1,2,3,5]
**Explanation:** After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).

**Example 2:**


**Input:** nums = [5,1,1,2,0,0]
**Output:** [0,0,1,1,2,5]
**Explanation:** Note that the values of nums are not necessairly unique.

**Constraints:**

* `1 <= nums.length <= 5 * 104`
* `-5 * 104 <= nums[i] <= 5 * 104`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/956104303/)

```go
func merge(left, right []int) []int {
    final := []int{}
    i := 0
    j := 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            final = append(final, left[i])
            i++
        } else {
            final = append(final, right[j])
            j++
        }
    }

    for ; i<len(left); i++ {
        final = append(final, left[i])
    }
    for ; j<len(right); j++ {
        final = append(final, right[j])
    }

    return final
}

func mergeSort(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }

    left := mergeSort(nums[:len(nums)/2])
    right := mergeSort(nums[len(nums)/2:])

    return merge(left, right)
}

func sortArray(nums []int) []int {
    return mergeSort(nums)
}
```