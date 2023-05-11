# [Problem statement](https://leetcode.com/problems/word-pattern)

Given a `pattern` and a string `s`, find if `s` follows the same pattern.

Here **follow** means a full match, such that there is a bijection between a letter in `pattern` and a **non-empty** word in `s`.

**Example 1:**


**Input:** pattern = "abba", s = "dog cat cat dog"
**Output:** true

**Example 2:**


**Input:** pattern = "abba", s = "dog cat cat fish"
**Output:** false

**Example 3:**


**Input:** pattern = "aaaa", s = "dog cat cat dog"
**Output:** false

**Constraints:**

* `1 <= pattern.length <= 300`
* `pattern` contains only lower-case English letters.
* `1 <= s.length <= 3000`
* `s` contains only lowercase English letters and spaces `' '`.
* `s` **does not contain** any leading or trailing spaces.
* All the words in `s` are separated by a **single space**.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/948496830/)

```go
func wordPattern(pattern string, s string) bool {
    lists := strings.Split(s, " ")
    if len(pattern) != len(lists) {
        return false
    }
    
    sr := map[string]byte{}
    rs := map[byte]string{}

    for idx, list := range lists {
        if val, ok := sr[list]; ok && val != pattern[idx] {
            return false
        }
        if val, ok := rs[pattern[idx]]; ok && val != list {
            return false
        }
        sr[list] = pattern[idx]
        rs[pattern[idx]] = list
    }

    return true
}
```