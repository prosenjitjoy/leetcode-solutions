func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    freq := make([][]int, len(nums)+1)

    for _, v := range nums {
        count[v]++;
    }

    for k, v := range count {
        freq[v] = append(freq[v], k)
    }

    var res []int
    for i:=len(freq)-1; i>0; i-- {
        for _, v := range freq[i] {
            res = append(res, v)
            if len(res) == k {
                return res
            }
        }
    }

    return res
}