func gridGame(grid [][]int) int64 {
    n := len(grid[0])
    prefixSumRow1 := make([]int, n)
    prefixSumRow2 := make([]int, n)
    
    prefixSumRow1[0] = grid[0][0]
    prefixSumRow2[0] = grid[1][0]

    for i:=1; i<n; i++ {
        prefixSumRow1[i] = grid[0][i] + prefixSumRow1[i-1]
        prefixSumRow2[i] = grid[1][i] + prefixSumRow2[i-1]
    }

    var res int64 = math.MaxInt64
    for i:=0; i<n; i++ {
        top := prefixSumRow1[n-1] - prefixSumRow1[i]
        var bottom int
        if i>0 {
            bottom = prefixSumRow2[i-1]
        }

        secondRobot := max(top, bottom)
        res = min(res, int64(secondRobot))
    }

    return res
}