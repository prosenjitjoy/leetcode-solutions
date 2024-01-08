func maxProfit(prices []int) int {
    left := 0
    right := 1
    max := 0
    profit := 0

    for i :=1; i<len(prices); i++ {
        right = i
        if prices[left]>prices[right] {
            left = right
        } else {
            profit = prices[right] - prices[left]
            if profit > max {
                max = profit
            }
        }
    }

    return max
}