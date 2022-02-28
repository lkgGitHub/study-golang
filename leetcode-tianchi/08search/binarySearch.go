package main

/*
360. 二分查找
难度：简单
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。


示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1


提示：
你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。
*/

func search(nums []int, target int) int {
	if len(nums) < 1 || nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func mySearch(nums []int, target int) int {
	if len(nums) < 1 || nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}
	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}

	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	for {

		if target == nums[mid] {
			return mid
		}
		println(left, mid, right)
		if left == right || left+1 == right {
			return -1
		}

		switch true {
		case target < nums[mid]:
			right, mid = mid, (mid+left)/2
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			left, mid = mid, (mid+right)/2
		}

	}

}
