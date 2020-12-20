package sorting

func CountSort(nums []int, max int) []int {
	// Build a slice of counts where the indicies are the
	// numbers and the values are the counts.
	counts := make([]int, max+1) // max(nums)
	results := make([]int, 0)
	for _, num := range nums {
		counts[num] += 1
	}
	// Walk through the counts and add the number to the results
	// as many times as it was present in the counts
	i := len(counts) - 1
	for i >= 0 {
		count := counts[i]
		for j := 0; j < count; j++ {
			results = append(results, i)
		}
		i -= 1
	}
	return results
}
