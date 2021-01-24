package dynamicprogramming

type Cake struct {
	Weight int
	Value  int
}

func MaxDuffelBagValue(cake_tuples []Cake, weight_capacity int) int {
	dp := make([]int, weight_capacity+1)
	for capacity := 0; capacity < weight_capacity+1; capacity++ {
		for i := 0; i < len(cake_tuples); i++ {
			cake := cake_tuples[i]
			weight, value := cake.Weight, cake.Value
			if weight <= capacity {
				remaining_weight := capacity - weight
				// Max with cake is the max value for the remaining weight
				// plus the current value
				max_with_cake := dp[remaining_weight] + value
				dp[capacity] = max(dp[capacity], max_with_cake)
			}
		}
	}
	return dp[weight_capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
