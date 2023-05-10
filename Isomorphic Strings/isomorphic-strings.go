func isIsomorphic(s string, t string) bool {
    m1 := map[byte]byte{}
    m2 := map[byte]byte{}

    for i:=0; i<len(s); i++ {
        v1, ok1 := m1[s[i]]
        if !ok1 {
            m1[s[i]] = t[i]
        } else if v1 != t[i] {
            return false
        }

        v2, ok2 := m2[t[i]]
        if !ok2 {
            m2[t[i]] = s[i]
        } else if v2 != s[i] {
            return false
        }
    }
    return true
}