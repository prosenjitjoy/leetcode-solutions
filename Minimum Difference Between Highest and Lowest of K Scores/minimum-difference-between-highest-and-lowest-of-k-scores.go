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