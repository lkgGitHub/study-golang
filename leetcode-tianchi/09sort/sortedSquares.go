package main

import "sort"

func sortedSquares(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	if len(nums) == 1 {
		nums[0] = nums[0] * nums[0]
		return nums
	}

	nums1 := make([]int, 0, len(nums)) // 递减序列
	nums2 := make([]int, 0, len(nums)) // 递增序列

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			nums1 = append(nums1, nums[i]*nums[i])
		} else {
			nums2 = append(nums2, nums[i]*nums[i])
		}
	}

	for i, j, n := len(nums1)-1, 0, 0; n < len(nums); n++ {
		if i < 0 {
			nums[n] = nums2[j]
			j++
			continue
		}
		if j >= len(nums2) {
			nums[n] = nums1[i]
			i--
			continue
		}

		if nums1[i] > nums2[j] {
			nums[n] = nums2[j]
			j++
		} else {
			nums[n] = nums1[i]
			i--
		}
	}
	return nums
}

// 直接排序
func sortedSquares1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}
