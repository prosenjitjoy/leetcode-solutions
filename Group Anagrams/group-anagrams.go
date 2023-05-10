func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]string{}
    for _, str := range strs {
        v := [26]int{}
        for i:=0; i<len(str); i++ {
            v[str[i] - 'a']++
        }
        m[v] = append(m[v], str)
    }
    res := [][]string{}
    for _, val := range m {
        res = append(res, val)
    }
    return res
}