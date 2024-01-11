func maxProduct(s string) int {
    palindrome := map[int]int{}

    for mask:=1; mask<(1<<len(s)); mask++ {
        subSequence := ""
        for i:=0; i<len(s); i++ {
            if mask & (1 << i) != 0 {
                subSequence += string(s[i])
            }
        }
        
        if subSequence == reverse(subSequence) {
            palindrome[mask] = len(subSequence)
        }
    }

    max := 0
    for k1, v1 := range palindrome {
        for k2, v2 := range palindrome {
            if k1 & k2 == 0 {
                if v1 * v2 > max {
                    max = v1 * v2
                }
            }
        }
    }

    return max
}

func reverse(s string) string {
    byteString := []rune(s)
    for i,j := 0, len(byteString)-1; i<j; i,j = i+1, j-1 {
        byteString[i], byteString[j] = byteString[j], byteString[i]
    }
    return string(byteString)
}