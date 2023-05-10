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