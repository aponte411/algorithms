package maths

import "testing"

func TestPower(t *testing.T) {
	res := Power(2.00000, 10)
	exp := 1024.00000
	if res != exp {
		t.Errorf("Expected %v, got %v", exp, res)
	}
	res2 := PingalaPower(2.00000, 10)
	if res2 != exp {
		t.Errorf("Expected %v, got %v", exp, res2)
	}
}
