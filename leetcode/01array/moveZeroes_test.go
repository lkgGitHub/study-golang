package array

import (
	"fmt"
	"testing"
)

/*
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数
*/
func moveZeroes(nums []int) {
	for i, j := 0, 1; j < len(nums) && i < len(nums); {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i] = nums[j]
			nums[j] = 0
			i++
			j++
		}
		if nums[i] != 0 {
			i++
		}
		j++
	}

}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 1, 0, 0, 12}, []int{1, 12, 0, 0, 0}},
		{[]int{1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0}},
	}
	for _, tt := range tests {
		moveZeroes(tt.nums)

		if len(tt.nums) != len(tt.expected) {
			t.Error(fmt.Sprintf("moveZeroes failed. expected: %v , actual: %v", tt.expected, tt.nums))
		}
		isPass := true
		for i := 0; i < len(tt.nums); i++ {
			if tt.expected[i] != tt.nums[i] {
				isPass = false
			}
		}
		if !isPass {
			t.Error(fmt.Sprintf("moveZeroes failed. expected: %v , actual: %v", tt.expected, tt.nums))
		}
	}
}
