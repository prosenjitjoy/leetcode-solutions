func bestClosingTime(customers string) int {
    n := len(customers)+1
    prefix := 0
    postfix := 0
    prefix_no := make([]int, n)
    postfix_yes := make([]int, n)

    for i, j := 1, n-2; i<n; i, j = i+1, j-1 {
        prefix_no[i] = prefix
        if customers[i-1] == 'N' {
            prefix++
        }

        if customers[j] == 'Y' {
            postfix++
        }
        postfix_yes[j] = postfix

    }

    min := math.MaxInt
    j := 0
    for i:=0; i<n; i++ {
        penalty := prefix_no[i] + postfix_yes[i]
        if penalty < min {
            min = penalty
            j = i
        }
    }

    return j
}