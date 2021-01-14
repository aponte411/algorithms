package searching

/*
We have a list of integers that are in range 1..n and has
length of n + 1. Find the integer that appears more than once
in our list
*/
func FindRepeat(nums []int) int {
	floor, ceiling := 1, len(nums)-1
	for floor < ceiling {
		pivot := floor + int((ceiling-floor)/2)
		// Split into two buckets
		lower_floor, lower_ceiling := floor, pivot
		upper_floor, upper_ceiling := pivot+1, ceiling
		// Count the number of numbers in lower bucket
		inLower := 0
		for _, num := range nums {
			if num >= lower_floor && num <= lower_ceiling {
				inLower += 1
			}
		}
		// Count distinct numbers in lower bucket
		distinctNums := lower_ceiling - lower_floor + 1
		// If the number of numbers in the lower bucket is greater
		// than the number of distinct numbers, then its in the lower bucket
		if inLower > distinctNums {
			floor, ceiling = lower_floor, lower_ceiling
		} else {
			floor, ceiling = upper_floor, upper_ceiling
		}
	}
	return floor
}
