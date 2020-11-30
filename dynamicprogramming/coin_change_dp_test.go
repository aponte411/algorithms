package dynamicprogramming

import "testing"

func TestCoinChangeDP(t *testing.T) {
	ans := CoinChangeDP(5, []int{1, 2, 5})
	if ans != 4 {
		t.Errorf("Expected 4, got %d", ans)
	}
}
