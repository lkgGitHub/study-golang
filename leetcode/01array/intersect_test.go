package array

import (
	"fmt"
	"testing"
)

/*
两个数组的交集 II
给定两个数组，编写一个函数来计算它们的交集。

https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

说明：
输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。

## 进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果nums1的大小比nums2小很多，哪种方法更优？
如果nums2的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/

// 哈希表法
// 时间复杂度：O(m+n)，其中 mm 和 nn 分别是两个数组的长度。需要遍历两个数组并对哈希表进行操作，哈希表操作的时间复杂度是 O(1)O(1)，因此总时间复杂度与两个数组的长度和呈线性关系。
// 空间复杂度：O(min(m,n))，其中 mm 和 nn 分别是两个数组的长度。对较短的数组进行哈希表的操作，哈希表的大小不会超过较短的数组的长度。为返回值创建一个数组 intersection，其长度为较短的数组的长度
func intersect(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]int, len(nums1))
	for _, n := range nums1 {
		if _, ok := numsMap[n]; ok {
			numsMap[n] = numsMap[n] + 1
		} else {
			numsMap[n] = 1
		}
	}

	nums := make([]int, 0, len(nums2))
	for _, n := range nums2 {
		if val, ok := numsMap[n]; ok {
			if val > 0 {
				nums = append(nums, n)
				numsMap[n] = val - 1
			}
		}
	}

	return nums
}

// todo 方法二：排序 + 双指针

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		expect []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}
	for _, tt := range tests {
		nums := intersect(tt.nums1, tt.nums2)
		if len(nums) != len(tt.expect) {
			t.Error(fmt.Sprintf("intersect failed. expect: %v , actual: %v", tt.expect, nums))
		}
		isPass := true
		numsMap := make(map[int]struct{}, len(nums))
		for _, n := range nums {
			numsMap[n] = struct{}{}
		}
		for _, n := range tt.expect {
			if _, ok := numsMap[n]; !ok {
				isPass = false
			}
		}
		if !isPass {
			t.Error(fmt.Sprintf("intersect failed. expect: %v , actual: %v", tt.expect, nums))
		}
	}
}
