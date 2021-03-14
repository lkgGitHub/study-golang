package array

import (
	"fmt"
	"testing"
)

/*
两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

提示：

2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
*/
// 哈希表
func twoSum(nums []int, target int) []int {
	tableMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if val, ok := tableMap[target-num]; ok {
			return []int{val, i}
		} else {
			tableMap[num] = i
		}
	}
	return nil
}

// 不优雅
func twoSum2(nums []int, target int) []int {
	var targetNum int
	if target%2 == 0 {
		targetNum = target / 2
	}

	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if val, ok := numsMap[num]; ok {
			if num == targetNum {
				return []int{val, i}
			}
		} else {
			numsMap[num] = i
		}
	}

	for i, num := range nums {
		if val, ok := numsMap[target-num]; ok && num != targetNum {
			return []int{i, val}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
	}
	for _, tt := range tests {
		actual := twoSum(tt.nums, tt.target)

		if len(actual) != len(tt.expected) {
			t.Error(fmt.Sprintf("twoSum failed. expected: %v , actual: %v", tt.expected, actual))
			return
		}
		isPass := true
		for i := 0; i < len(tt.expected); i++ {
			if tt.expected[i] != actual[i] {
				isPass = false
			}
		}
		if !isPass {
			t.Error(fmt.Sprintf("twoSum failed. expected: %v , actual: %v", tt.expected, actual))
		}
	}
}
