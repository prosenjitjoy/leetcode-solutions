func partitionString(s string) int {
    seen := [26]bool{}
    count := 1

    for _, v := range s {
        if seen[v-'a'] {
            count++
            seen = [26]bool{}
        }

        seen[v-'a'] = true
    }

    return count
}