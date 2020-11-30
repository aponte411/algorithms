package permutations

func PrintStringPermutations(str string, i, n int) {
	// Base case
	if i == n-1 {
		println(str)
		return
	}
	// Go through string
	for j := i; j < n; j++ {
		// swap character at index i with current
		// character
		str = swap(str, i, j)
		// recurse  through next character
		PrintStringPermutations(str, i+1, n)
		// backtrack
		str = swap(str, i, j)
	}
}

// Helper function to swap i-th with j-th byte
func swap(s string, i, j int) string {
	var res []byte
	for k := 0; k < len(s); k++ {
		if k == i {
			res = append(res, s[j])
		} else if k == j {
			res = append(res, s[i])
		} else {
			res = append(res, s[k])
		}
	}
	return string(res)
}
