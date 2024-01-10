# [Problem statement](https://leetcode.com/problems/unique-length-3-palindromic-subsequences)

Given a string `s`, return _the number of **unique palindromes of length three** that are a **subsequence** of_ `s`.

Note that even if there are multiple ways to obtain the same subsequence, it is still only counted **once**.

A **palindrome** is a string that reads the same forwards and backwards.

A **subsequence** of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

* For example, `"ace"` is a subsequence of `"abcde"`.

**Example 1:**


**Input:** s = "aabca"
**Output:** 3
**Explanation:** The 3 palindromic subsequences of length 3 are:
- "aba" (subsequence of "aabca")
- "aaa" (subsequence of "aabca")
- "aca" (subsequence of "aabca")

**Example 2:**


**Input:** s = "adc"
**Output:** 0
**Explanation:** There are no palindromic subsequences of length 3 in "adc".

**Example 3:**


**Input:** s = "bbcbaba"
**Output:** 4
**Explanation:** The 4 palindromic subsequences of length 3 are:
- "bbb" (subsequence of "bbcbaba")
- "bcb" (subsequence of "bbcbaba")
- "bab" (subsequence of "bbcbaba")
- "aba" (subsequence of "bbcbaba")

**Constraints:**

* `3 <= s.length <= 105`
* `s` consists of only lowercase English letters.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1142379339/)

```go
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
```