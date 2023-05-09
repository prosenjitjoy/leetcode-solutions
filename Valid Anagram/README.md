# [Problem statement](https://leetcode.com/problems/valid-anagram)

Given two strings `s` and `t`, return `true` _if_ `t` _is an anagram of_ `s`_, and_ `false` _otherwise_.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

**Example 1:**

**Input:** s = "anagram", t = "nagaram"
**Output:** true

**Example 2:**

**Input:** s = "rat", t = "car"
**Output:** false

**Constraints:**

* `1 <= s.length, t.length <= 5 * 104`
* `s` and `t` consist of lowercase English letters.

**Follow up:** What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

<br />

# [Solution in go](https://leetcode.com/submissions/detail/947244759/)

```go
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m1 := map[rune]int{}
    m2 := map[rune]int{}
    for _, val := range s {
        m1[val]++
    }
    for _, val := range t {
        m2[val]++
    }

    for _, val := range s {
        if m1[val] != m2[val] {
            return false
        }
    }
    return true
}
```