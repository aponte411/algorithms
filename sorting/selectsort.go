package sorting

// O(n^2) time, O(1) space
func SelectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		// Find min index
		minIndex := findMinIndex(nums, i)
		// Swap current element at index i with element at minIndex
		swap(nums, i, minIndex)
	}
}

func findMinIndex(nums []int, start int) int {
	min := start
	for i := start + 1; i < len(nums); i++ {
		if nums[i] < nums[min] {
			min = i
		}
	}
	return min
}

func swap(nums []int, p1, p2 int) {
	nums[p1], nums[p2] = nums[p2], nums[p1]
}
