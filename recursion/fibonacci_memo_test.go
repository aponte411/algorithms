package recursion

import "testing"

func TestFibMemo(t *testing.T) {
	ans := FibonacciMemo(10)
	if ans != 55 {
		t.Errorf("Expected 55, got %d", ans)
	}
}
