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