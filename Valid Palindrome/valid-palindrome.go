func isPalindrome(s string) bool {
    var str []rune
    for _, v := range s {
        if v >= 'a' && v <= 'z' || v >= '0' && v <= '9' {
            str = append(str, v)
        } else if v >= 'A' && v <= 'Z' {
            str = append(str, v-'A'+'a')
        }
    }

    for i,j := 0, len(str)-1; i<j; i,j = i+1, j-1 {
        if str[i] != str[j] {
            return false
        }
    }

    return true
}