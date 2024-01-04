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

                _, ok1 := row[i][board[i][j]]
                _, ok2 := col[j][board[i][j]]
                _, ok3 := box[k][board[i][j]]

                if ok1 || ok2 || ok3 {
                    return false
                } else {
                    row[i][board[i][j]] = true
                    col[j][board[i][j]] = true
                    box[k][board[i][j]] = true
                }
            }
        }
    } 
    return true
}