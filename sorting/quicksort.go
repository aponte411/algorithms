package sorting

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, low, high int) {
	if low < high {
		pivot := partition(nums, low, high)
		quickSort(nums, low, pivot-1)
		quickSort(nums, pivot+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]
	for low < high {
		for nums[high] >= pivot && low < high {
			high -= 1
		}
		// swap
		nums[low], nums[high] = nums[high], nums[low]
		for nums[low] <= pivot && low < high {
			low += 1
		}
		nums[low], nums[high] = nums[high], nums[low]
	}
	return low
}
