package array

import (
	"testing"
)

// 27. 移除元素 https://leetcode.cn/problems/remove-element/

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"示例1", args{[]int{3, 2, 2, 3}, 3}, 2},
		{"示例2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}

// todo 双指针法
