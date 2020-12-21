package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 5, 10, 90, 100}
	res := BinarySearch(nums, 90)
	if res != 3 {
		t.Errorf("Expected %v, found %v", 3, res)
	}
}
