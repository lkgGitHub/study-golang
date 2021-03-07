package array

import (
	"fmt"
	"sort"
	"testing"
)

// 217. 存在重复元素
/*
给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。



示例 1:

输入: [1,2,3,1]
输出: true
示例 2:

输入: [1,2,3,4]
输出: false
示例 3:

输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
*/

// 方法二：哈希表
// 时间复杂度：O(N)O(N)，其中 NN 为数组的长度。
// 空间复杂度：O(N)O(N)，其中 NN 为数组的长度。
func containsDuplicate2(nums []int) bool {
	numMap := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := numMap[n]; ok {
			return true
		} else {
			numMap[n] = struct{}{}
		}
	}
	return false
}

// 方法一：排序
// 时间复杂度：O(N*logN)，其中 NN 为数组的长度。需要对数组进行排序。
// 空间复杂度：O(logN)，其中 NN 为数组的长度。注意我们在这里应当考虑递归调用栈的深度。
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		expect bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, tt := range tests {
		if tt.expect != containsDuplicate(tt.nums) {
			t.Error("containsDuplicate failed")
		}
	}
}
