package searching

import "testing"

func TestFindDuplicates(t *testing.T) {
	nums := []int{1, 2, 5, 5, 5, 5}
	res1 := FindDuplicate(nums)
	res2 := FindDuplicateV2(nums)
	if res1 != 5 && res2 != 5 {
		t.Errorf("Expected 5, got %v", res1)
		t.Errorf("Expected 5, got %v", res2)
	}
}
