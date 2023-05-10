# [Problem statement](https://leetcode.com/problems/pascals-triangle)

Given an integer `numRows`, return the first numRows of **Pascal's triangle**.

In **Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:

![](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif) 

**Example 1:**

**Input:** numRows = 5
**Output:** [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

**Example 2:**

**Input:** numRows = 1
**Output:** [[1]]

**Constraints:**

* `1 <= numRows <= 30`

<br />

# [Solution in go](https://leetcode.com/submissions/detail/947644578/)

```go
func generate(numRows int) [][]int {
    tri := [][]int{[]int{1}}
    for i:=1; i<numRows; i++ {
        tmp := make([]int, i+1)
        tmp[0] = 1
        tmp[i] = 1
        for j:=1; j<i; j++ {
            tmp[j] = tri[i-1][j-1] + tri[i-1][j]
        }
        tri = append(tri, tmp)
    }
    return tri
}
```