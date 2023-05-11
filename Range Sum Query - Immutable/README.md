# [Problem statement](https://leetcode.com/problems/range-sum-query-immutable)

Given an integer array `nums`, handle multiple queries of the following type:

1. Calculate the **sum** of the elements of `nums` between indices `left` and `right` **inclusive** where `left <= right`.

Implement the `NumArray` class:

* `NumArray(int[] nums)` Initializes the object with the integer array `nums`.
* `int sumRange(int left, int right)` Returns the **sum** of the elements of `nums` between indices `left` and `right` **inclusive** (i.e. `nums[left] + nums[left + 1] + ... + nums[right]`).

**Example 1:**


**Input**
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
**Output**
[null, 1, -1, -3]

**Explanation**
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

**Constraints:**

* `1 <= nums.length <= 104`
* `-105 <= nums[i] <= 105`
* `0 <= left <= right < nums.length`
* At most `104` calls will be made to `sumRange`.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/948397959/)

```go
type NumArray struct {
    psum []int
}

func Constructor(nums []int) NumArray {
    psum := make([]int, len(nums))
    psum[0] = nums[0]

    for i:=1; i<len(nums); i++ {
        psum[i] = psum[i-1] + nums[i]
    }
    fmt.Println(psum)
    return NumArray{
        psum: psum,
    }
}

func (this *NumArray) SumRange(left int, right int) int {
    if left - 1 < 0 {
        return this.psum[right]
    }
    return this.psum[right] - this.psum[left-1]
}
```