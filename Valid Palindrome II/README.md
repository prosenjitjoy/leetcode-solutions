# [Problem statement](https://leetcode.com/problems/valid-palindrome-ii)

Given a string `s`, return `true` _if the_ `s` _can be palindrome after deleting **at most one** character from it_.

**Example 1:**


**Input:** s = "aba"
**Output:** true

**Example 2:**


**Input:** s = "abca"
**Output:** true
**Explanation:** You could delete the character 'c'.

**Example 3:**


**Input:** s = "abc"
**Output:** false

**Constraints:**

* `1 <= s.length <= 105`
* `s` consists of lowercase English letters.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1188530437/)

```go
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
```