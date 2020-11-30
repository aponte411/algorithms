package recursion

import "testing"

func TestFibonacci(t *testing.T) {
	ans := Fibonacci(10)
	if ans != 55 {
		t.Errorf("Wanted 55, got %d", ans)
	}
}
