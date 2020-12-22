package searching

// O(log n) time, O(1) space
func FindRotationPoint(words []string) int {
	left, right := 0, len(words)-1
	first_word := words[0]
	for left < right {
		pivot := left + int((right-left)/2)
		// If current guess (pivot) comes after first word, then the rotation point
		// must be to the right of the pivot
		if words[pivot] >= first_word {
			left = pivot
		} else {
			// If its before our first word, then rotation point
			// must be to the left of the pivot
			right = pivot
		}
		// if left and right converge, the left item is the last item
		// of the rotated list, making the right item the rotation point
		if left+1 == right {
			return right
		}
	}
	return -1
}
