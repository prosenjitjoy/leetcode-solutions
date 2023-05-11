func wordPattern(pattern string, s string) bool {
    lists := strings.Split(s, " ")
    if len(pattern) != len(lists) {
        return false
    }
    
    sr := map[string]byte{}
    rs := map[byte]string{}

    for idx, list := range lists {
        if val, ok := sr[list]; ok && val != pattern[idx] {
            return false
        }
        if val, ok := rs[pattern[idx]]; ok && val != list {
            return false
        }
        sr[list] = pattern[idx]
        rs[pattern[idx]] = list
    }

    return true
}