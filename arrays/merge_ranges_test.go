package arrays

import (
	"reflect"
	"testing"
)

func TestMergeRanges(t *testing.T) {
	list1 := [][]int{
		{1, 3},
		{2, 4},
	}
	ans1 := MergeRanges(list1)
	exp1 := [][]int{{1, 4}}
	if !reflect.DeepEqual(ans1, exp1) {
		t.Errorf("Expected %v got %v", exp1, ans1)
	}

	list2 := [][]int{
		{5, 6},
		{6, 8},
	}
	ans2 := MergeRanges(list2)
	exp2 := [][]int{{5, 8}}
	if !reflect.DeepEqual(ans2, exp2) {
		t.Errorf("Expected %v got %v", exp2, ans2)
	}

}
