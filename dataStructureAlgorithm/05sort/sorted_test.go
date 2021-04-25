package main

import (
	"fmt"
	"sort"
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
	if len(nums) <= 1 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		// 注意：相对于nums[j+1] = nums[j]，已经又走了一论，现在的j比最后一个这行表达式里的j小1，所以再加1才是现在的j
		nums[j+1] = value
	}

	return nums
}

// 选择排序(写着就感觉比冒泡和插入排序复杂）
func selectionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums)-1; i++ {
		min := nums[i]
		a := i
		b := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				b = j
			}
		}
		nums[a], nums[b] = nums[b], nums[a]
	}
	return nums
}

func quicksort2(nums []int) {

}

func quickSort(nums []int, low, height int) {
	if height > low {
		return
	}
	index := partition(nums, low, height)
	quickSort(nums, low, index-1)
	quickSort(nums, index+1, height)
}

func partition(nums []int, low int, height int) int {

}

func goSort(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func TestName(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{3, 4, 5, 2, 1}, []int{1, 2, 3, 4, 5}},
		//{[]int{4, 5, 6, 1, 3, 2}, []int{1, 2, 3, 4, 5, 6}},
		//{[]int{1}, []int{1}},
		//{[]int{}, []int{}},
	}
	for _, tt := range tests {
		quickSort(tt.nums, 0, len(tt.nums))
		actual := tt.nums
		isPass := true
		if len(actual) != len(tt.expected) {
			t.Errorf("sorted failed. expect: %v, actual: %v", tt.expected, actual)
			continue
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

func TestDemo(t *testing.T) {

	a := []int{1, 2, 3}
	fmt.Println(a[0:2])
	//n := 10
	//var depth int
	//for i := n; i > 0; i >>= 1 {
	//	depth++
	//}
	//println(depth * 2)
}
