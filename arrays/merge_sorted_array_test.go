package arrays

import (
	"reflect"
	"testing"
)

func TestMergeArrays(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	exp := []int{1, 2, 2, 3, 5, 6}
	m := 3
	n := 3
	MergeArraysV1(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, exp) {
		t.Errorf("Expected %v, got %v", nums1, exp)
	}
	nums3 := []int{2, 4, 6, 8}
	nums4 := []int{1, 7}
	exp2 := []int{1, 2, 4, 6, 7, 8}
	res := MergeArraysV2(nums3, nums4)
	if !reflect.DeepEqual(res, exp2) {
		t.Errorf("Expected %v, got %v", exp2, res)
	}
}
