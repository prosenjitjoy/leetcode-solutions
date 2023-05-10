# [Problem statement](https://leetcode.com/problems/longest-common-prefix)

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string `""`.

**Example 1:**


**Input:** strs = ["flower","flow","flight"]
**Output:** "fl"

**Example 2:**


**Input:** strs = ["dog","racecar","car"]
**Output:** ""
**Explanation:** There is no common prefix among the input strings.

**Constraints:**

* `1 <= strs.length <= 200`
* `0 <= strs[i].length <= 200`
* `strs[i]` consists of only lowercase English letters.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/947619270/)

```go
func longestCommonPrefix(strs []string) string {
    pre := ""
    for i:=0; i<len(strs[0]); i++ {
        tmp := strs[0][i]
        for j:=0; j<len(strs); j++ {
            if i >= len(strs[j]) || strs[j][i] != tmp {
                return pre
            }
        }
        pre += string(tmp)
    }
    return pre
}
```