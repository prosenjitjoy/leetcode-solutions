func gridGame(grid [][]int) int64 {
    n := len(grid[0])
    for i:=1; i<n; i++ {
        grid[0][i] += grid[0][i-1]
        grid[1][i] += grid[1][i-1]
    }

    var res int64 = math.MaxInt64
    for i:=0; i<n; i++ {
        top := grid[0][n-1] - grid[0][i]
        var bottom int
        if i>0 {
            bottom = grid[1][i-1]
        }

        secondRobot := max(top, bottom)
        res = min(res, int64(secondRobot))
    }

    return res
}