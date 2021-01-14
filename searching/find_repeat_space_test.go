package searching

import (
	"testing"
)

func TestFindRepeat(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	exp := 2
	res := FindRepeat(nums)
	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
