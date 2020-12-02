package hashtables

import "testing"

func TestTwoSum(t *testing.T) {
	ans1 := TwoSum([]int{2, 4}, 1)
	ans2 := TwoSum([]int{2, 4}, 6)
	ans3 := TwoSum([]int{1, 2, 3, 4, 5, 6}, 7)
	if ans1 {
		t.Error("Expected false, got true")
	}
	if !ans2 {
		t.Error("Expected true, got false")
	}
	if !ans3 {
		t.Error("Expected true, got false")
	}

}
