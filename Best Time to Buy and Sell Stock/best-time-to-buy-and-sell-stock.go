func maxProfit(prices []int) int {
    buy := 0
    max := 0

    for sell := 1; sell<len(prices); sell++ {
        if prices[buy]>prices[sell] {
            buy = sell
        } else if prices[sell] - prices[buy] > max {
            max = prices[sell] - prices[buy]
        }
    }

    return max
}