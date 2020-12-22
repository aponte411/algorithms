package sorting

import (
	"testing"

	"reflect"
)

func TestBubbleSort(t *testing.T) {
	input := []int{5, 2, 3, 1}
	exp := []int{1, 2, 3, 5}
	BubbleSort(input)
	if !reflect.DeepEqual(exp, input) {
		t.Errorf("Expected %v, got %v", exp, input)
	}
}
