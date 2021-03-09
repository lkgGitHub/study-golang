package array

import "testing"

/*
只出现一次的数字

tags: 位运算  哈希表
（整数包括0）
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
*/

// 位运算, 异或运算：相同为0，不同为1. 异或同一个数两次，原数不变。
// 时间复杂度：O(n)，其中 n 是数组长度。只需要对数组遍历一次。
// 空间复杂度：O(1)。
func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] = nums[0] ^ nums[i]
	}
	return nums[0]
}

// 哈希表
func singleNumber2(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for _, n := range nums {
		if _, ok := numMap[n]; ok {
			numMap[n] = 2
		} else {
			numMap[n] = 1
		}
	}

	for k, v := range numMap {
		if v == 1 {
			return k
		}
	}
	return 0
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}
	for _, tt := range tests {
		if tt.expect != singleNumber(tt.nums) {
			t.Errorf("singleNumber failed. expect: %d, actual: %d", tt.expect, tt.nums)
		}
	}
}

// 异或运算 XOR： 相同为0，不同为1. 异或同一个数两次，原数不变。
func TestXOR(t *testing.T) {
	println("0 ^ 0: ", 0^0)
	println("0 ^ 1: ", 0^1)
	println("1 ^ 1: ", 1^1)
	println("0 ^ 2: ", 0^2)
	println("1 ^ 2: ", 1^2)
	println("2 ^ 1: ", 2^1)
	println("2 ^ 2: ", 2^2)
	println("2 ^ 3: ", 2^3)
}
