package dynamicprogramming

import "testing"

func TestFibonacci(t *testing.T) {
	res := Fibonacci(10)
	if res != 55 {
		t.Errorf("Expected 55, got %v", res)
	}
}
