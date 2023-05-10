func canPlaceFlowers(flowerbed []int, n int) bool {
    bed := []int{0}
    bed = append(bed, flowerbed...)
    bed = append(bed, 0)
    count := 0

    for i:=1; i<len(bed)-1; i++ {
        if bed[i] == 0 && bed[i-1] == 0 && bed[i+1] == 0 {
            count++
            bed[i] = 1
        } 
    }
    return count >= n
}