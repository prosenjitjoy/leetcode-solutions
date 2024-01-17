func findRepeatedDnaSequences(s string) []string {
    freq := map[string]int{}
    res := []string{}

    for i:=0; i<=len(s)-10; i++ {
        freq[s[i:i+10]]++
    }

    for k, v := range freq {
        if v > 1 {
            res = append(res, k)
        }
    }

    return res
}