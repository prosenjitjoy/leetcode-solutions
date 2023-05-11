func findDisappearedNumbers(nums []int) []int {
    ans := []int{}
    for _, num := range nums {
        nums[abs(num)-1] = abs(nums[abs(num)-1]) * -1
    }
    for idx, num := range nums {
        if num > 0 {
            ans = append(ans, idx+1)
        }
    }
    return ans
}

func abs(num int) int {
    if num < 0 { 
        return num * -1
    } else {
        return num
    }
}