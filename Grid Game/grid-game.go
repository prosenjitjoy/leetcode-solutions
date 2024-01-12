func gridGame(grid [][]int) int64 {
    for i:=1; i<len(grid[0]); i++ {
        grid[0][i] += grid[0][i-1]
        grid[1][i] += grid[1][i-1]
    }

    var res int64 = math.MaxInt64
    for i:=0; i<len(grid[0]); i++ {
        top := grid[0][len(grid[0])-1] - grid[0][i]
        var bottom int
        if i>0 {
            bottom = grid[1][i-1]
        }

        secondRobot := max(top, bottom)
        res = min(res, int64(secondRobot))
    }

    return res
}