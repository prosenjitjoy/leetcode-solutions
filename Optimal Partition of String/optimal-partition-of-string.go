func partitionString(s string) int {
    curSet := map[rune]bool{}
    res := 1

    for _, v := range s {
        if curSet[v] {
            res++
            curSet = map[rune]bool{}
        }

        curSet[v] = true
    }

    return res
}