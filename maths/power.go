package maths

// O(n) time, O(1) space
func Power(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1 / x
		N = -N
	}
	ans := 1.0
	for i := 0; i < N; i++ {
		ans *= x
	}
	return ans
}

// O(log n) time, and O(1) space
// PingalaPower Aglorithm, goes back to 2nd century BCE
func PingalaPower(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1 / x
		N = -N
	}
	return power(x, N)
}

// T(n) <= T(n/2) + 2 -> O(log n)
func power(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
    // Reducing space by half
	half := power(x, n/2)
	if n % 2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
