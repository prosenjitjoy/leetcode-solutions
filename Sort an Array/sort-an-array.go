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

    for i<len(left) {
        final = append(final, left[i])
        i++
    }
    for j<len(right) {
        final = append(final, right[j])
        j++
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