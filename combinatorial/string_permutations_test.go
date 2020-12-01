package combinatorial

import "testing"

func TestStringPermutations(t *testing.T) {
	str := "ABC"
	PrintStringPermutations(str, 0, len(str))
	t.Log(str)
}
