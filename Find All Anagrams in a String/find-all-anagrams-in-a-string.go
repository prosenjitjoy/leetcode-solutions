func findAnagrams(s string, p string) []int {
    res := []int{}
    sLen := len(s)
    pLen := len(p)

    if pLen > sLen {
        return res
    }

    sCount := make([]int, 26)
    pCount := make([]int, 26)

    for i:=0; i<pLen; i++ {
        pCount[p[i]-'a']++
        sCount[s[i]-'a']++
    }

    left := 0
    if slices.Equal(pCount, sCount) {
        res = append(res, left)
    }
    sCount[s[left]-'a']--
    left++

    for right:=pLen; right<sLen; right++ {
        sCount[s[right]-'a']++

        if slices.Equal(pCount, sCount) {
            res = append(res, left)
        }
        sCount[s[left]-'a']--
        left++
    }

    return res
}