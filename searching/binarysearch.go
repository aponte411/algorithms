package searching

/*
1. The value at a given index tells us a lot about what's to the left and what's to the right.
2. We don't have to look at every item in the list. By inspecting the middle item, we can "rule out" half of the list.
3. We can use this approach over and over, cutting the problem in half until we have the answer. This is sometimes called "divide and conquer."
*/

// O(log n) time, O(1) space
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// Guess
		pivot := int((left + right) / 2)
		// If our guess is our  target were done
		if nums[pivot] == target {
			return pivot
			// If our guess is less than our target, we have a new floor
			// so search through the right
		} else if nums[pivot] < target {
			left = pivot + 1
			// If our guess is greater than our target, we have a new ceiling
			// so search through the left
		} else {
			right = pivot - 1
		}
	}
	return -1
}
