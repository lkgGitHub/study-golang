package main

import (
	"fmt"
	_ "net/http/pprof"
	"testing"
)

// 循环
func climbStairs2(n int) int {
	return 1
}

// 避免重复
var stairMap = make(map[int]int)

// 避免堆栈溢出，设置最大递归深度
var depth uint

// 递归
func climbStairs(n int) int {
	depth++
	if depth > 100 {
		panic("depth > 100")
	}

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if climb, ok := stairMap[n]; ok {
		return climb
	} else {
		var nums int
		nums += climbStairs(n-1) + climbStairs(n-2)
		stairMap[n] = nums
		return nums
	}
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{45, 1836311903},
	}
	for _, tt := range tests {
		actual := climbStairs(tt.n)
		if tt.expected != actual {
			t.Error(fmt.Sprintf("TestRotate failed. expect: %v , actual: %v", tt.expected, actual))
		}
	}
}
