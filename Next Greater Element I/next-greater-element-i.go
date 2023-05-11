func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m := map[int]int{}
    arr := []int{}

    for i, v := range nums2 {
        m[v] = i
    }
    for _, num := range nums1 {
        found := false
        for i:=m[num]+1; i<len(nums2); i++ {
            if nums2[i] > num {
                arr = append(arr, nums2[i])
                found = true
                break
            }
        }
        if !found {
            arr = append(arr, -1)
        }
    }
    return arr
}