func leastBricks(wall [][]int) int {
    gap := map[int]int{}

    for i := 0; i<len(wall); i++ {
        cum := 0
        for j:=0; j<len(wall[i])-1; j++ {
            cum += wall[i][j]
            gap[cum]++
        }
    }

    max := 0
    for _, v := range gap {
        if v > max {
            max = v
        }
    }

    return len(wall) - max
}