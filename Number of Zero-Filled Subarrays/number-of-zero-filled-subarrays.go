func zeroFilledSubarray(nums []int) int64 {
    var res int64
    var count int64

    for _, v := range nums {
        if v == 0 {
            count++
        } else {
            count = 0
        }

        res += count
    }

    return res
}