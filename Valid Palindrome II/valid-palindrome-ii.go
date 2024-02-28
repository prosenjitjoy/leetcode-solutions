func validPalindrome(s string) bool {
    for i,j := 0,len(s)-1; i<j; i,j = i+1,j-1 {
        if s[i] != s[j] {
            return isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1)
        }
    }
    return true
}

func isPalindrome(s string, start int, end int) bool {
    for i,j := start,end; i<j; i,j = i+1,j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}