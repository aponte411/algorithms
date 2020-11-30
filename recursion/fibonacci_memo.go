package recursion

func FibonacciMemo(n int) int {
	memo := make(map[int]int)
	return fib(n, memo)
}

func fib(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}
