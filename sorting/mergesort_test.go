package sorting

import (
	"testing"

	"reflect"
)

func TestMergeSort(t *testing.T) {
	input := []int{5, 2, 3, 1}
	exp := []int{1, 2, 3, 5}
	res := MergeSort(input)
    if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
