package arrays

import "testing"

func TestMergeArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	MergeArrays(nums1, m, nums2, n)
	if nums1[len(nums1)-1] != 6 {
		t.Errorf("Expected 6, got %v", nums1[len(nums1)])
	}
}
