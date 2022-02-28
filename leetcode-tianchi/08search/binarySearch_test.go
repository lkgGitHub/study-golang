package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"示例 1:", args{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{"示例 2:", args{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
		{"示例 3:", args{[]int{5}, 5}, 0},
		{"示例 4:", args{[]int{-1, 0, 3, 5, 9, 12}, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
