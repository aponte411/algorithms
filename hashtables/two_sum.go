package hashtables

/*
Are there two numbers that sum up to a target?

Input: [2, 4] target=1
Output: False

Input: [2, 4] target=6
Output: True

Input: [1,2,3,4,5,6] target=7
Output: True
*/
type IntSet map[int]bool

func TwoSum(nums []int, target int) bool {
	seen := make(IntSet)
	for _, num := range nums {
		remainder := target - num
		if _, ok := seen[remainder]; ok {
			return true
		} else {
			seen[num] = true
		}
	}
	return false
}
