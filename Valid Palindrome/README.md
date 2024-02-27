# [Problem statement](https://leetcode.com/problems/valid-palindrome)

A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` _if it is a **palindrome**, or_ `false` _otherwise_.

**Example 1:**


**Input:** s = "A man, a plan, a canal: Panama"
**Output:** true
**Explanation:** "amanaplanacanalpanama" is a palindrome.

**Example 2:**


**Input:** s = "race a car"
**Output:** false
**Explanation:** "raceacar" is not a palindrome.

**Example 3:**


**Input:** s = " "
**Output:** true
**Explanation:** s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

**Constraints:**

* `1 <= s.length <= 2 * 105`
* `s` consists only of printable ASCII characters.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1187622852/)

```go
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
```