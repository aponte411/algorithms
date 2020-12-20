package sorting

import "testing"

func TestCountSort(t *testing.T) {
	result := CountSort([]int{37, 89, 41, 65, 91, 53}, 100)
	if result[0] != 91 {
		t.Errorf("Expected %v, got %v", 91, result[0])
	}
}
