func interchangeableRectangles(rectangles [][]int) int64 {
    ratio := map[float64]int64{}
    for _, rectangle := range rectangles {
        ratio[float64(rectangle[0])/float64(rectangle[1])]++
    }

    var res int64
    for _, v := range ratio {
        if v > 0 {
            res += (v * (v-1)) / 2
        }
    }

    return res
}
