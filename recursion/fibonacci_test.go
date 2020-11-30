package recursion

import "testing"

func TestFibonacci(t *testing.T) {
	ans1 := Fibonacci(10)
	ans2 := FibonacciMemo(10)
	if ans1 != 55 {
		t.Errorf("Wanted 55, got %d", ans1)
	}
	if ans2 != 55 {
		t.Errorf("FibonacciMemo() expected 55, got %d", ans2)
	}
}
