func findAnagrams(s string, p string) []int {
    var res []int
    if len(p) > len(s) {
        return res
    }

    sCount := map[byte]int{}
    pCount := map[byte]int{}

    for i:=0; i<len(p); i++ {
        pCount[p[i]]++
        sCount[s[i]]++
    }

    left := 0
    if isEqual(pCount, sCount) {
        res = append(res, left)
    }

    for right:=len(p); right<len(s); right++ {
        sCount[s[right]]++
        sCount[s[left]]--

        if sCount[s[left]] == 0 {
            delete(sCount, s[left])
        }

        left++
        if isEqual(pCount, sCount) {
            res = append(res, left)
        }
    }

    return res
}

func isEqual(m1, m2 map[byte]int) bool {
    if len(m1) != len(m2) {
        return false
    }

    for k, v := range m1 {
        if w, ok := m2[k]; !ok || w != v {
            return false
        }
    }

    return true
}