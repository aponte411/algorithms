package dynamicprogramming

import (
	"testing"
)

func TestCakeThief(t *testing.T) {
	exp := 4
	cake_tuples := []Cake{
		Cake{Weight: 2, Value: 1},
	}
	res := MaxDuffelBagValue(cake_tuples, 9)
	if res != exp {
		t.Errorf("Expected %v got %v", res, exp)
	}
}
