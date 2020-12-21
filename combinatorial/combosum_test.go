package combinatorial

import (
	"reflect"
	"testing"
)

func TestComboSum(t *testing.T) {
	exp := [][]int{
		{3, 2, 2},
		{7},
	}
	inputs := []int{2, 3, 6, 7}
	res := ComboSum(inputs, 7)
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
