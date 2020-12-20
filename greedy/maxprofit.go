package greedy

import "math"

func MaxProfit(prices []int) int {
    minPrice := math.MaxInt64
    maxProfit := 0
    for _, price := range prices {
        minPrice = min(minPrice, price)
        maxProfit = max(maxProfit, price - minPrice)
    }
    return maxProfit
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
