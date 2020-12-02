package hashtables

/*
Check whether any permutation of an input string is a palindrome

Input: civic
Output: true

Input: ivicc
Output: true
*/

type StringSet map[string]bool

func HasPalindromePermutation(str string) bool {
	odds := make(StringSet)
	for _, ch := range str {
		if _, ok := odds[string(ch)]; ok {
			delete(odds, string(ch))
		} else {
			odds[string(ch)] = true
		}
	}
	return len(odds) <= 1
}
