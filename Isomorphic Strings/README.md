# [Problem statement](https://leetcode.com/problems/isomorphic-strings)

Given two strings `s` and `t`, _determine if they are isomorphic_.

Two strings `s` and `t` are isomorphic if the characters in `s` can be replaced to get `t`.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

**Example 1:**

**Input:** s = "egg", t = "add"
**Output:** true

**Example 2:**

**Input:** s = "foo", t = "bar"
**Output:** false

**Example 3:**

**Input:** s = "paper", t = "title"
**Output:** true

**Constraints:**

* `1 <= s.length <= 5 * 104`
* `t.length == s.length`
* `s` and `t` consist of any valid ascii character.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/947678857/)

```go
func isIsomorphic(s string, t string) bool {
    m1 := map[byte]byte{}
    m2 := map[byte]byte{}

    for i:=0; i<len(s); i++ {
        v1, ok1 := m1[s[i]]
        if !ok1 {
            m1[s[i]] = t[i]
        } else if v1 != t[i] {
            return false
        }

        v2, ok2 := m2[t[i]]
        if !ok2 {
            m2[t[i]] = s[i]
        } else if v2 != s[i] {
            return false
        }
    }
    return true
}
```