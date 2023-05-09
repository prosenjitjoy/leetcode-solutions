func getConcatenation(nums []int) []int {
    ans := []int{}
    ans = append(ans, nums...)
    ans = append(ans, nums...)
    return ans
}