package main

import (
	"testing"
)

// 冒泡排序（Bubble Sort）
func bubbleSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				flag = true // 表示有数据交换
			}
		}
		if !flag {
			break
		}
	}
	return nums
}

// 插入排序（Insertion Sort）
func insertionSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		flag := false
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		if flag {

		}

	}
	return nums
}

func TestName(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		actual := bubbleSort(tt.nums)
		isPass := true
		if len(actual) != len(tt.expected) {
			panic("")
		}
		for i := 0; i < len(actual); i++ {
			if tt.expected[i] != actual[i] {
				isPass = false
				break
			}
		}
		if !isPass {
			t.Errorf("sorted failed. expect: %v, actual: %v", tt.expected, actual)
		}
	}
}
