func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m := map[int]int{}
    res := make([]int, len(nums1))
    st := []int{}

    for idx, num := range nums1 {
        m[num] = idx
        res[idx] = -1
    }

    for _, num := range nums2 {
        for len(st)>0 && num > st[len(st)-1] {
            res[m[st[len(st)-1]]] = num
            st = st[:len(st)-1]
        }
        if _, ok := m[num]; ok {
            st = append(st, num)
        }
    }

    return res
}