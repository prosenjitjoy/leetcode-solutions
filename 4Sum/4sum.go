func fourSum(nums []int, target int) [][]int {
    u := New(nums)
    u.kSum(4, 0, target)
    return u.GetResult()
}

type uniqueSum struct {
    res [][]int
    quad []int
    nums []int
}

func New(nums []int) uniqueSum {
    return uniqueSum{
        res: [][]int{},
        quad: []int{},
        nums: mergeSort(nums),
    }
}

func (u *uniqueSum) kSum(k int, start int, target int) {
    if k != 2 {
        for i:=start; i<=len(u.nums)-k; i++ {
            if i>start && u.nums[i] == u.nums[i-1] {
                continue
            }
            u.quad = append(u.quad, u.nums[i])
            u.kSum(k-1, i+1, target-u.nums[i])
            u.quad = u.quad[:len(u.quad)-1]
        }
        return
    }

    l, r := start, len(u.nums)-1
    for l<r {
        if u.nums[l]+u.nums[r] < target {
            l++
        } else if u.nums[l]+u.nums[r] > target {
            r--
        } else {
            u.res = append(u.res, slices.Concat(u.quad, []int{u.nums[l], u.nums[r]}))
            l++
            for l<r && u.nums[l] == u.nums[l-1] {
                l++
            }
        }
    }
}

func (u *uniqueSum) GetResult() [][]int {
    return u.res
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
    var final []int
    var i, j int

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