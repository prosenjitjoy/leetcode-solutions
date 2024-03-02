# [Problem statement](https://leetcode.com/problems/reverse-string)

Write a function that reverses a string. The input string is given as an array of characters `s`.

You must do this by modifying the input array [in-place](https://en.wikipedia.org/wiki/In-place%5Falgorithm) with `O(1)` extra memory.

**Example 1:**

**Input:** s = ["h","e","l","l","o"]
**Output:** ["o","l","l","e","h"]

**Example 2:**

**Input:** s = ["H","a","n","n","a","h"]
**Output:** ["h","a","n","n","a","H"]

**Constraints:**

* `1 <= s.length <= 105`
* `s[i]` is a [printable ascii character](https://en.wikipedia.org/wiki/ASCII#Printable%5Fcharacters).

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1191455211/)

```go
func reverseString(s []byte)  {
    for i,j:=0,len(s)-1; i<j; i,j=i+1,j-1 {
        s[i],s[j] = s[j],s[i]
    }
}
```