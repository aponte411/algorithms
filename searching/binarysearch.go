package searching

// O(log n) time, O(1) space
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		pivot := int((left + right) / 2)
		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] < nums[left] {
			left = left - 1
		} else {
			right = right + 1
		}
	}
	return -1
}
