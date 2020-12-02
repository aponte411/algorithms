package arrays

import "testing"

func TestMergeRanges(t *testing.T) {
	ans1 := MergeRanges([]int{1, 3}, []int{2, 4})
	exp1 := []int{[]int{1, 4}}
	ans2 := MergeRanges([]int{5, 6}, []int{6, 8})
	exp2 := []int{[]int{5, 8}}

}
