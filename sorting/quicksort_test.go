package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{5, 2, 1, 3}
	exp := []int{1, 2, 3, 5}
	QuickSort(input)
	if !reflect.DeepEqual(exp, input) {
		t.Errorf("Expected %v, got %v", exp, input)
	}
}
