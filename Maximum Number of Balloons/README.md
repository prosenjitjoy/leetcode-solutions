# [Problem statement](https://leetcode.com/problems/maximum-number-of-balloons)

Given a string `text`, you want to use the characters of `text` to form as many instances of the word **"balloon"** as possible.

You can use each character in `text` **at most once**. Return the maximum number of instances that can be formed.

**Example 1:**

**![](https://assets.leetcode.com/uploads/2019/09/05/1536_ex1_upd.JPG)**


**Input:** text = "nlaebolko"
**Output:** 1

**Example 2:**

**![](https://assets.leetcode.com/uploads/2019/09/05/1536_ex2_upd.JPG)**


**Input:** text = "loonbalxballpoon"
**Output:** 2

**Example 3:**


**Input:** text = "leetcode"
**Output:** 0

**Constraints:**

* `1 <= text.length <= 104`
* `text` consists of lower case English letters only.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/948442383/)

```go
func maxNumberOfBalloons(text string) int {
    match := "balloon"
    freq := map[rune]int{}
    for _, v := range text {
        freq[v]++
    }
    m := map[rune]int{}
    for _, v := range match {
        m[v]++
    }
    min := len(text)
    for k, v := range m {
        val, ok := freq[k]
        if !ok {
            return 0
        }
        if val / v < min {
            min = val / v
        }
    }
    return min
}

```