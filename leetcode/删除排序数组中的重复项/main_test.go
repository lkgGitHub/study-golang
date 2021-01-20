/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/
package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums   []int
		length int
	}{
		{[]int{}, 0},
		{[]int{1, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 2}, 2},
		{[]int{1, 2, 3}, 3},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, tt := range tests {
		if tt.length != removeDuplicates(tt.nums) {
			t.Error(tt.nums)
		}
	}

}

// 方法：双指针法
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[length] = nums[i]
			length++
		}
	}
	print("[")
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", nums[i])
	}
	println("]")
	return length
}
