# [Problem statement](https://leetcode.com/problems/group-anagrams)

Given an array of strings `strs`, group **the anagrams** together. You can return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

**Example 1:**

**Input:** strs = ["eat","tea","tan","ate","nat","bat"]
**Output:** [["bat"],["nat","tan"],["ate","eat","tea"]]

**Example 2:**

**Input:** strs = [""]
**Output:** [[""]]

**Example 3:**

**Input:** strs = ["a"]
**Output:** [["a"]]

**Constraints:**

* `1 <= strs.length <= 104`
* `0 <= strs[i].length <= 100`
* `strs[i]` consists of lowercase English letters.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/947633083/)

```go
func groupAnagrams(strs []string) [][]string {
    m := map[[26]int][]string{}
    for _, str := range strs {
        v := [26]int{}
        for i:=0; i<len(str); i++ {
            v[str[i] - 'a']++
        }
        m[v] = append(m[v], str)
    }
    res := [][]string{}
    for _, val := range m {
        res = append(res, val)
    }
    return res
}
```