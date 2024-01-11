func maxProduct(s string) int {
    palindrome := map[int]int{}

    for mask:=1; mask<(1<<len(s)); mask++ {
        subSequence := ""
        for i:=0; i<len(s); i++ {
            if mask & (1 << i) != 0 {
                subSequence += string(s[i])
            }
        }
        
        if isPalindrome(subSequence) {
            palindrome[mask] = len(subSequence)
        }
    }

    res := 0
    for k1, v1 := range palindrome {
        for k2, v2 := range palindrome {
            if k1 & k2 == 0 {
                res = max(res, v1 * v2)
            }
        }
    }

    return res
}

func isPalindrome(s string) bool {
    for i,j := 0, len(s)-1; i<j; i,j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}