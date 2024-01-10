func minSwaps(s string) int {
    max := 0
    extraClosingBracket := 0

    for _, v := range s {
        if v == ']' {
            extraClosingBracket++
        } else {
            extraClosingBracket--
        }

        if extraClosingBracket > max {
            max = extraClosingBracket
        }
    }

    return (max+1)/2
}