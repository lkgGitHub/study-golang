package leetcode_tianchi

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"示例1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"示例2,数组中同一个元素在答案里不能重复出现", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"示例3", args{[]int{3, 3}, 6}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		m[n] = i
	}
	for i, n := range nums {
		if j, ok := m[target-n]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}
