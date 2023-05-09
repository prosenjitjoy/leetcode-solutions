func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m1 := map[rune]int{}
    m2 := map[rune]int{}
    for _, val := range s {
        m1[val]++
    }
    for _, val := range t {
        m2[val]++
    }

    for _, val := range s {
        if m1[val] != m2[val] {
            return false
        }
    }
    return true
}