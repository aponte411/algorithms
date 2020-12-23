package searching

import "sort"

// O(n), O(n) space
func FindDuplicateV1(nums []int) int {
	seen := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = true
		} else {
			return num
		}
	}
	return -1
}

// O(n logn) time, O(1) space but "destroys" original input
func FindDuplicateV2(nums []int) int {
	// Sorting in place gives us O(1) space
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == nums[right] {
			return nums[right]
		}
		left++
		right--
	}
	return -1
}

// O(n logn) time, O(1) space
func FindDuplicate(nums []int) int {
	floor, ceiling := 1, len(nums)-1
	for floor < ceiling {
		// Split into two ranges of numbers
		midpoint := floor + int((ceiling-floor)/2)
		lower_floor, lower_ceiling := floor, midpoint
		upper_floor, upper_ceiling := midpoint+1, ceiling
		// Count number of numbers in first range (lower)
		inLower := 0
		for _, num := range nums {
			if lower_floor <= num && num <= lower_ceiling {
				inLower += 1
			}
		}
		// Get distinct number of numbers in lower range
		distinctLower := lower_ceiling - lower_floor + 1
		if inLower > distinctLower {
			// Theres a duplicate in the lower range, use the same
			// approach in the new smaller range
			floor, ceiling = lower_floor, lower_ceiling
		} else {
			// The duplicate is in the upper range
			floor, ceiling = upper_floor, upper_ceiling
		}
	}
	// once were done we arrive at the duplicate value
	return floor
}
