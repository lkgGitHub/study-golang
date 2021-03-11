package array

import (
	"fmt"
	"testing"
)

/*
加一

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。


示例1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。

示例2：
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。

示例 3：
输入：digits = [0]
输出：[1]

提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9

*/
// 数组
func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 < 10 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	nums := make([]int, len(digits)+1, len(digits)+1)
	nums[0] = 1
	_ = copy(nums[1:], digits)
	return nums
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits   []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{0}, []int{1}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		actual := plusOne(tt.digits)
		isPass := true
		for i := 0; i < len(actual); i++ {
			if actual[i] != tt.expected[i] {
				isPass = false
			}
		}
		if !isPass {
			t.Error(fmt.Sprintf("plusOne failed. expected: %v , actual: %v", tt.expected, actual))
		}
	}
}

func TestPlusOneDemo(t *testing.T) {
	digits := []int{2, 3}
	nums := make([]int, len(digits)+1, len(digits)+1)
	nums[0] = 1
	_ = copy(nums[1:], digits)
	fmt.Println(nums)
}
