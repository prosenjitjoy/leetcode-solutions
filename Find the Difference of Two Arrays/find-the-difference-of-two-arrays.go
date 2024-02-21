func findDifference(nums1 []int, nums2 []int) [][]int {
    numSet1 := map[int]bool{}
    numSet2 := map[int]bool{}

    for i := range nums1 {
        numSet1[nums1[i]] = true
    }
    for i := range nums2 {
        numSet2[nums2[i]] = true
    }

    ans1 := []int{}
    ans2 := []int{}

    for key := range numSet1 {
        if !numSet2[key] {
            ans1 = append(ans1, key)
        }
    }
    for key := range numSet2 {
        if !numSet1[key] {
            ans2 = append(ans2, key)
        }
    }

    return [][]int{ans1, ans2}
}