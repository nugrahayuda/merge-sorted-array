package mergesortedarray

import "slices"

// Leetcode 88 https://leetcode.com/problems/merge-sorted-array/

func Merge(nums1 []int, m int, nums2 []int, n int) []int {

	j := 0
	for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[j]
		j++
	}

	slices.Sort(nums1)

	return nums1
}
