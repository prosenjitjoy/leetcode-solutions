func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m := map[rune]int{}
    for _, val := range s {
        m[val]++
    }

    for _, val := range t {
        if n, ok := m[val]; !ok || n == 0 {
            return false
        }
        m[val]--
    }

    return true
}