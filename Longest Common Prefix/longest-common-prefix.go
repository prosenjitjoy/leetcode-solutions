func longestCommonPrefix(strs []string) string {
    pre := ""
    for i:=0; i<len(strs[0]); i++ {
        tmp := strs[0][i]
        for j:=0; j<len(strs); j++ {
            if i >= len(strs[j]) || strs[j][i] != tmp {
                return pre
            }
        }
        pre += string(tmp)
    }
    return pre
}