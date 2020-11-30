package recursion

import "testing"

func TestCoinChange(t *testing.T) {
	ans1 := CoinChange(4, []int{1, 2, 3})
	ans2 := CoinChangeMemo(4, []int{1, 2, 3})
	if ans1 != 4 {
		t.Errorf("CoinChange() expected 4, got %d", ans1)
	}
	if ans2 != 4 {
		t.Errorf("CoinChangeMemo() expected 4, got %d", ans2)
	}
}
