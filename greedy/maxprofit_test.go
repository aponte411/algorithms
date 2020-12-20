package greedy

import "testing"


func TestMaxProfit(t *testing.T) {
    res := MaxProfit([]int{7,1,5,3,6,4})
    if res != 5 {
        t.Errorf("Expected 5, got %v", res)
    }
}
