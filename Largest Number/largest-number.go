func largestNumber(nums []int) string {
    numStr := make([]string, len(nums))
    for i, v := range nums {
        numStr[i] = strconv.Itoa(v)
    }

    sort.Slice(numStr, func(a, b int) bool {
        return numStr[a] + numStr[b] > numStr[b] + numStr[a]
    })

    ansStr := strings.Join(numStr, "")

    if ansStr[0] == '0' {
        return "0"
    }
    return ansStr
}