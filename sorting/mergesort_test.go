package sorting

import "testing"

func TestMergeSort(t *testing.T){
    input := []int{5,2,3,1}
    exp := []int{1,2,3,5}
    res := MergeSort(input)
    if res[0] != exp[0] {
        t.Errorf("Expected %v, got %v", exp[0], res[0])
    }
}
