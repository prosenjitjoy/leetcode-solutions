func longestConsecutive(nums []int) int {
    set := map[int]bool{}
    for _, num := range nums {
        set[num] = true
    }
    
    max := 0
    freq := map[int]int{}
    for k := range set {
        if !set[k-1] {
            i := k
            for set[i] {
                freq[k]++
                i++
            }

            if freq[k] > max {
                max = freq[k]
            }
        }
    }

    return max
}