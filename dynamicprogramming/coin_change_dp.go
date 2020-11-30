package dynamicprogramming

func CoinChangeDP(amount int, denominations []int) int {
	solutions := make([]int, amount+1)
	solutions[0] = 1
	for _, coin := range denominations {
		for t := coin; t < amount+1; t++ {
			// move on to next coin if we cant use it
			if t-coin < 0 {
				continue
			}
			solutions[t] += solutions[t-coin]
		}
	}
	return solutions[amount]
}
