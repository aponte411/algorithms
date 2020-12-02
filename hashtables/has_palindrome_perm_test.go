package hashtables

import "testing"

func TestHasPalindromePerm(t *testing.T) {
	ans1 := HasPalindromePermutation("aabccbdd")
	if !ans1 {
		t.Error("Expected True got False")
	}

	ans2 := HasPalindromePermutation("aabcd")
	if ans2 {
		t.Error("Expected False, got true")
	}
}
