package main

import (
	"testing"
)

func sorted(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
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
		actual := sorted(tt.nums)
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
