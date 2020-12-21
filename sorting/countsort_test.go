package sorting

import (
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	res := CountSort([]int{37, 89, 41, 65, 91, 53}, 100)
	exp := []int{91, 89, 65, 53, 41, 37}
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
