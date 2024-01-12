# [Problem statement](https://leetcode.com/problems/find-all-anagrams-in-a-string)

Given two strings `s` and `p`, return _an array of all the start indices of_ `p`_'s anagrams in_ `s`. You may return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

**Example 1:**


**Input:** s = "cbaebabacd", p = "abc"
**Output:** [0,6]
**Explanation:**
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

**Example 2:**


**Input:** s = "abab", p = "ab"
**Output:** [0,1,2]
**Explanation:**
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".

**Constraints:**

* `1 <= s.length, p.length <= 3 * 104`
* `s` and `p` consist of lowercase English letters.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1144523798/)

```go
func findAnagrams(s string, p string) []int {
    res := []int{}
    sLen := len(s)
    pLen := len(p)

    if pLen > sLen {
        return res
    }

    sCount := make([]int, 26)
    pCount := make([]int, 26)

    for i:=0; i<pLen; i++ {
        pCount[p[i]-'a']++
        sCount[s[i]-'a']++
    }

    left := 0
    if slices.Equal(pCount, sCount) {
        res = append(res, left)
    }
    sCount[s[left]-'a']--
    left++

    for right:=pLen; right<sLen; right++ {
        sCount[s[right]-'a']++

        if slices.Equal(pCount, sCount) {
            res = append(res, left)
        }
        sCount[s[left]-'a']--
        left++
    }

    return res
}
```