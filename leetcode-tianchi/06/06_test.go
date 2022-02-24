package _6

import (
	"fmt"
	"testing"
)

/*
只出现一次的数字
难度：简单
收藏
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
*/
func TestSingleNumber(t *testing.T) {
	a := []int{2, 2, 1}
	println(singleNumber(a))

	b := []int{4, 1, 2, 1, 2}
	println(singleNumber(b))

	c := []int{2, 2, 1}
	println(singleNumber(c))
}

// 按位异或操作 ^
func singleNumber(nums []int) int {
	m := 0
	for i := 0; i < len(nums); i++ {
		m ^= nums[i]
	}
	return m
}

// 暴力法，使用额外空间
func singleNumber2(nums []int) int {
	m := make(map[int]struct{}, len(nums)/2)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
		} else {
			m[n] = struct{}{}
		}
	}
	for k := range m {
		return k
	}
	return 0
}

func TestSingleNumberIII(t *testing.T) {
	nums := []int{0, 0, 1, 2}

	fmt.Println(singleNumberIII(nums))
}
func singleNumberIII(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}

	n, m := nums[0], nums[1]
	p := n ^ m
	for i := 2; i < len(nums); i++ {
		p ^= nums[i]
		if p == n {
			m = nums[i]
		} else {
			n = nums[i]
		}
	}
	return []int{n, m}
}
