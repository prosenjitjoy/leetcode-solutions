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