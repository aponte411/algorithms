package arrays

import "sort"

// O(N) time, O(N) space
func MergeRanges(nums [][]int) [][]int {
    // Sort by first element, i.e. sort by start of meetings
    // to get natural order of meetings.
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	merged := make([][]int, 0)
	merged = append(merged, nums[0])
	for _, num := range nums {
		start, end := num[0], num[1]
		next_ending := merged[len(merged)-1]
        // If the last meetings end time is
        // greater than the start time of this
        // current meeting, merge.
		if next_ending[1] >= start {
            // assign to last element in merged to
            // always get the next ending.
			merged[len(merged)-1] = []int{
				next_ending[0],
				max(next_ending[1], end),
			}
		} else {
			merged = append(merged, []int{start, end})
		}
	}
	return merged

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
