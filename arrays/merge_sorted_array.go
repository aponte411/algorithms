package arrays

func MergeArraysV1(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j -= 1
		} else {
			nums1[k] = nums1[i]
			i -= 1
		}
		k -= 1
	}
	if j >= 0 {
		copy(nums1[:k+1], nums2[:j+1])
	}
}

// Alternative
func MergeArraysV2(left, right []int) []int {
	i, j, k := 0, 0, 0
	size := len(left) + len(right)
	merged := make([]int, size)
	for k < size {
		if i > len(left)-1 && j >= len(right)-1 {
			merged[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			merged[k] = left[i]
			i++
		} else if left[i] < right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
		k++
	}
	return merged
}
