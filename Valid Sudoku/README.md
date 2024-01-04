# [Problem statement](https://leetcode.com/problems/valid-sudoku)

Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

1. Each row must contain the digits `1-9` without repetition.
2. Each column must contain the digits `1-9` without repetition.
3. Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9` without repetition.

**Note:**

* A Sudoku board (partially filled) could be valid but is not necessarily solvable.
* Only the filled cells need to be validated according to the mentioned rules.

**Example 1:**

![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png) 


**Input:** board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
**Output:** true

**Example 2:**


**Input:** board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
**Output:** false
**Explanation:** Same as Example 1, except with the **5** in the top left corner being modified to **8**. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

**Constraints:**

* `board.length == 9`
* `board[i].length == 9`
* `board[i][j]` is a digit `1-9` or `'.'`.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1136819268/)

```go
func isValidSudoku(board [][]byte) bool {
    type pos struct {
        x byte
        y byte
    }
    row := [9]map[byte]bool{}
    col := [9]map[byte]bool{}
    box := map[pos]map[byte]bool{}

    for i := range board {
        for j := range board[i] {
            if board[i][j] == '.' {
                continue
            }

            k := pos{x: byte(i/3), y: byte(j/3)}
            if row[i] == nil {
                row[i] = map[byte]bool{}
            }
            if col[j] == nil {
                col[j] = map[byte]bool{}
            }
            if box[k] == nil {
                box[k] = map[byte]bool{}
            }

            if _, ok := row[i][board[i][j]]; ok {
                return false
            } else {
                row[i][board[i][j]] = true
            }

            if _, ok := col[j][board[i][j]]; ok {
                return false
            } else {
                col[j][board[i][j]] = true
            }

            if _, ok := box[k][board[i][j]]; ok {
                return false
            } else {
                box[k][board[i][j]] = true
            }
        }
    } 
    return true
}
```