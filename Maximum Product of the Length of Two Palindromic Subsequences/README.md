# [Problem statement](https://leetcode.com/problems/maximum-product-of-the-length-of-two-palindromic-subsequences)

Given a string `s`, find two **disjoint palindromic subsequences** of `s` such that the **product** of their lengths is **maximized**. The two subsequences are **disjoint** if they do not both pick a character at the same index.

Return _the **maximum** possible **product** of the lengths of the two palindromic subsequences_.

A **subsequence** is a string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters. A string is **palindromic** if it reads the same forward and backward.

**Example 1:**

![example-1](https://assets.leetcode.com/uploads/2021/08/24/two-palindromic-subsequences.png) 


**Input:** s = "leetcodecom"
**Output:** 9
**Explanation**: An optimal solution is to choose "ete" for the 1st subsequence and "cdc" for the 2nd subsequence.
The product of their lengths is: 3 * 3 = 9.

**Example 2:**


**Input:** s = "bb"
**Output:** 1
**Explanation**: An optimal solution is to choose "b" (the first character) for the 1st subsequence and "b" (the second character) for the 2nd subsequence.
The product of their lengths is: 1 * 1 = 1.

**Example 3:**


**Input:** s = "accbcaxxcxx"
**Output:** 25
**Explanation**: An optimal solution is to choose "accca" for the 1st subsequence and "xxcxx" for the 2nd subsequence.
The product of their lengths is: 5 * 5 = 25.

**Constraints:**

* `2 <= s.length <= 12`
* `s` consists of lowercase English letters only.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1143695628/)

```go
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

func isPalindrome(s []rune) bool {
    for i,j := 0, len(s)-1; i<j; i,j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}
```