package combinatorial

func ComboSum(nums []int, target int) [][]int {
	combos := make([][]int, 0)
	helper(nums, &combos, []int{}, 0, target, 0)
	return combos
}

// bottom up helper that finds all combos for a given candidate and stops
// when it reaches the target sum
func helper(nums []int, combos *[][]int, combo []int, curr, target, index int) {
	// If we've reached the target, add the current combination to the results
	if curr == target {
		*combos = append(*combos, combo)
		return
	}
	// If we pass it, return
	if curr > target {
		return
	}
	for i := index; i < len(nums); i++ {
		// Add current candidate then combo everything from current combo
		comboCopy := make([]int, 0)
		comboCopy = append(comboCopy, nums[i])
		comboCopy = append(comboCopy, combo...)
		// Increment current sum by candidate
		helper(nums, combos, comboCopy, curr+nums[i], target, i)
	}
}
