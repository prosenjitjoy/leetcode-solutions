func maxNumberOfBalloons(text string) int {
    match := "balloon"
    freq := map[rune]int{}
    for _, v := range text {
        freq[v]++
    }
    m := map[rune]int{}
    for _, v := range match {
        m[v]++
    }
    min := len(text)
    for k, v := range m {
        val, ok := freq[k]
        if !ok {
            return 0
        }
        if val / v < min {
            min = val / v
        }
    }
    return min
}
