package recursion

/*
Write a function that given:
     1. an amount of money
     2. a list of coin denominations
computes the number of ways to make the amount of money
with the coins available.
*/

func CoinChange(amount int, denominations []int) int {
	return numberOfWays(amount, denominations, 0)
}

func numberOfWays(curr int, denominations []int, index int) int {
	if curr == 0 {
		return 1
	}
	if curr < 0 {
		return 0
	}
	if index == len(denominations) {
		return 0
	}

	cand := denominations[index]
	var numWays int
	for curr >= 0 {
		numWays += numberOfWays(curr, denominations, index+1)
		curr -= cand
	}
	return numWays
}

type MemoKey struct {
	val   int
	index int
}

func CoinChangeMemo(amount int, denominations []int) int {
	memo := make(map[MemoKey]int)
	return numberOfWaysMemo(amount, denominations, 0, memo)
}

func numberOfWaysMemo(curr int, denominations []int, index int, memo map[MemoKey]int) int {
	key := MemoKey{
		val:   curr,
		index: index,
	}
	if _, ok := memo[key]; ok {
		return memo[key]
	}
	if curr == 0 {
		return 1
	}
	if curr < 0 {
		return 0
	}
	if index == len(denominations) {
		return 0
	}

	cand := denominations[index]
	var numWays int
	for curr >= 0 {
		numWays += numberOfWaysMemo(curr, denominations, index+1, memo)
		curr -= cand
	}
	memo[key] = numWays
	return numWays
}
