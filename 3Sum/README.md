# [Problem statement](https://leetcode.com/problems/3sum)

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

**Example 1:**


**Input:** nums = [-1,0,1,2,-1,-4]
**Output:** [[-1,-1,2],[-1,0,1]]
**Explanation:** 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

**Example 2:**


**Input:** nums = [0,1,1]
**Output:** []
**Explanation:** The only possible triplet does not sum up to 0.

**Example 3:**


**Input:** nums = [0,0,0]
**Output:** [[0,0,0]]
**Explanation:** The only possible triplet sums up to 0.

**Constraints:**

* `3 <= nums.length <= 3000`
* `-105 <= nums[i] <= 105`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1196661483/)

```go
func threeSum(nums []int) [][]int {
    res := [][]int{}
    nums = mergeSort(nums)

    for i, v := range nums {
        if i>0 && v == nums[i-1] {
            continue
        }

        l, r := i+1, len(nums)-1
        for l < r {
            sum := v + nums[l] + nums[r]
            if sum > 0 {
                r--
            } else if sum < 0 {
                l++
            } else {
                res = append(res, []int{v, nums[l], nums[r]})
                l++
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
            }
        }
    }

    return res
}

func mergeSort(items []int) []int {
    if len(items) < 2 {
        return items
    }

    left := mergeSort(items[:len(items)/2])
    right := mergeSort(items[len(items)/2:])
    return merge(left, right)
}

func merge(a []int, b []int) []int {
    final := []int{}
    i, j := 0, 0

    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }

    for i < len(a) {
        final = append(final, a[i])
        i++
    }

    for j < len(b) {
        final = append(final, b[j])
        j++
    }

    return final
}
```