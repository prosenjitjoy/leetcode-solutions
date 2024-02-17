type NumMatrix struct {
    MatrixSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    row := len(matrix)
    col := len(matrix[0])

    m := make([][]int, row+1)
    for i := range m {
        m[i] = make([]int, col+1)
    }

    for r:=0; r<row; r++ {
        prefix := 0
        for c:=0; c<col; c++ {
            prefix += matrix[r][c]
            above := m[r][c+1]
            m[r+1][c+1] = prefix + above
        }
    }

    return NumMatrix{MatrixSum: m}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    bottomRight := this.MatrixSum[row2+1][col2+1]
    above := this.MatrixSum[row1][col2+1]
    left := this.MatrixSum[row2+1][col1]
    topLeft := this.MatrixSum[row1][col1]

    return bottomRight - above - left + topLeft
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */