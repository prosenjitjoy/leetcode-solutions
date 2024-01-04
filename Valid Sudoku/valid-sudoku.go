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
            if board[i][j] != '.' {
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
    } 
    return true
}