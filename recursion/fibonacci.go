package recursion

// O(2^n) exponential time, O(1) space
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// O(n) time, O(n) space
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

// O(n) time, O(1) space
func FibIterative(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	prev_prev := 0
	prev := 1
	var curr int
	for i := 0; i < n-1; i++ {
		curr = prev_prev + prev
		prev_prev = prev
		prev = curr
	}
	return curr
}
