package main

import "testing"

/*
	查找第一个值等于给定值的元素，包含重复的数据
*/

// 循环实现：序数据集合中不存在重复的数据
// 注意1. 循环退出条件 low<=high
// 注意2. mid=(low+high)/2 可能溢出。位运算：low+((high-low)>>1)
// 注意3. low 和 high 的更新： low=mid+1，high=mid-1
func binarySearch11(nums []int, value int) int {
	if len(nums) == 0 {
		return -1
	}

	low := 0
	height := len(nums) - 1

	for low <= height {
		mid := low + ((height - low) >> 1)
		if nums[mid] == value {
			if mid == 0 {
				return mid
			}
			for mid > 0 {
				if nums[mid-1] == value {
					mid--
				} else {
					break
				}
			}
			return mid

		} else if value < nums[mid] {
			height = mid - 1
		} else if nums[mid] < value {
			low = mid + 1
		}
	}

	return -1
}

// 递归实现：序数据集合中不存在重复的数据
func binarySearch1(nums []int, value int) int {
	return binarySearchRecursion(nums, 0, len(nums)-1, value)
}
func binarySearchRecursion1(nums []int, low, height, value int) int {
	if len(nums) == 0 || low > height {
		return -1
	}
	mid := low + ((height - low) >> 1)

	if nums[mid] == value {
		return mid
	} else if value < nums[mid] {
		return binarySearchRecursion(nums, low, mid-1, value)
	} else {
		return binarySearchRecursion(nums, mid+1, height, value)
	}
}

func TestBinarySearch1(t *testing.T) {
	tests := []struct {
		nums     []int
		value    int
		expected int
	}{
		{[]int{}, 0, -1},
		{[]int{2, 2, 2}, 2, 0},
		{[]int{1, 2, 2, 4, 5, 6, 7}, 2, 1},
		//{[]int{1, 2, 3, 4, 5, 6, 7}, 2, 1},
		//{[]int{1, 2, 3, 4, 5, 6, 7}, 3, 2},
		//{[]int{1, 2, 3, 4, 5, 6, 7}, 4, 3},
		//{[]int{1, 2, 3, 4, 5, 6, 7}, 5, 4},
		//{[]int{1, 2, 3, 4, 5, 6, 7}, 6, 5},
		//{[]int{1, 2, 3, 4, 5, 6, 7}, 7, 6},
		//{[]int{1, 2, 3, 4, 5, 6, 7}, 8, -1},
	}
	for _, tt := range tests {
		actual := binarySearch11(tt.nums, tt.value)
		if tt.expected != actual {
			t.Errorf("sorted failed. expect: %v, actual: %v", tt.expected, actual)
			continue
		}

	}
}
