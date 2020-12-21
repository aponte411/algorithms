package arrays

import (
	"reflect"
	"testing"
)

func TestMergeArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	exp := []int{1, 2, 2, 3, 5, 6}
	m := 3
	n := 3
	MergeArrays(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, exp) {
		t.Errorf("Expected %v, got %v", nums1, exp)
	}
}
