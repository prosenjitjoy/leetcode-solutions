# [Problem statement](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string)

Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

**Example 1:**


**Input:** haystack = "sadbutsad", needle = "sad"
**Output:** 0
**Explanation:** "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

**Example 2:**


**Input:** haystack = "leetcode", needle = "leeto"
**Output:** -1
**Explanation:** "leeto" did not occur in "leetcode", so we return -1.

**Constraints:**

* `1 <= haystack.length, needle.length <= 104`
* `haystack` and `needle` consist of only lowercase English characters.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1145038587/)

```go
func strStr(haystack string, needle string) int {
    N := len(needle)
    LPS := make([]int, N)

    for prevLPS, i := 0, 1; i<N; {
        if needle[i] == needle[prevLPS] {
            LPS[i] = prevLPS + 1
            prevLPS++
            i++
        } else {
            if prevLPS == 0 {
                LPS[i] = 0
                i++
            } else {
                prevLPS = LPS[prevLPS - 1]
            }
        }
    }

    for i, j := 0, 0; i<len(haystack); {
        if haystack[i] == needle[j] {
            i++
            j++
        } else {
            if j == 0 {
                i++
            } else {
                j = LPS[j-1]
            }
        }

        if j == N {
            return i - N
        }
    }

    return -1
}
```