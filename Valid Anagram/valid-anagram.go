func isAnagram(s string, t string) bool {
    m1 := map[rune]int{}
    m2 := map[rune]int{}
    for _, val := range s {
        m1[val]++
    }
    for _, val := range t {
        m2[val]++
    }
    if len(m1) != len(m2) {
        return false
    }
    for _, val := range s {
        if m1[val] != m2[val] {
            return false
        }
    }
    return true
}