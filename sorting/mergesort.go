package sorting

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := int(len(nums) / 2)
	left := MergeSort(nums[:pivot])
	right := MergeSort(nums[pivot:])
	return merge(left, right)
}

// O(n) time, O(n) space
func merge(left, right []int) []int {
	i, j, k := 0, 0, 0
	size := len(left) + len(right)
	res := make([]int, size)
	for k < size {
		if i > len(left)-1 && j >= len(right)-1 {
			res[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			res[k] = left[i]
			i++
		} else if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	return res
}
