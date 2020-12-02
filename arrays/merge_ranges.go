package arrays

import "sort"

func MergeRanges(nums [][]int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	merged := make([]int, nums[len(nums)-1])
	for _, num := range nums {
		start, end := num
		next_ending := merged[len(merged)-1]
		if next_ending[1] >= start {
			merged[len(merged)-1] = []int{
				next_ending[0], max(next_ending[1], end),
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
