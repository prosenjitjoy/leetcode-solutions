func countPalindromicSubsequence(s string) int {
    type palindrome struct {
        mid byte
        sides byte
    }
    res := map[palindrome]bool{}
    left := map[byte]bool{}
    right := map[byte]int{}

    for i:=0; i<len(s); i++ {
        right[s[i]-'a']++
    }

    for i:=0; i<len(s); i++ {
        right[s[i]-'a']--
        if right[s[i]-'a'] == 0 {
            delete(right, s[i]-'a')
        }

        for j:=0; j<26; j++ {
            if _, ok := right[byte(j)]; ok && left[byte(j)] {
                res[palindrome{mid: s[i]-'a', sides: byte(j)}] = true
            }
        }

        left[s[i]-'a'] = true
    }
    
    return len(res)
}