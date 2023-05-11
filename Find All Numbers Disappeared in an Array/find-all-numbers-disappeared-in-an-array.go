func findDisappearedNumbers(nums []int) []int {
    ans := []int{}
    found := make([]bool, len(nums))

    for _, num := range nums {
        found[num-1] = true
    }
    for i, v := range found {
        if v == false {
            ans = append(ans, i+1)
        }
    }
    return ans
}