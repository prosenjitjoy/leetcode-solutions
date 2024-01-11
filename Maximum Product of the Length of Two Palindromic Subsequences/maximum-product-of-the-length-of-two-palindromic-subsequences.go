func maxProduct(s string) int {
    palindrome := map[int]int{}

    for mask:=1; mask<(1<<len(s)); mask++ {
        var subSequence []rune
        for i, v := range s {
            if mask & (1 << i) != 0 {
                subSequence = append(subSequence, v)
            }
        }
        
        if isPalindrome(subSequence) {
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

func isPalindrome(s []rune) bool {
    for i,j := 0, len(s)-1; i<j; i,j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}