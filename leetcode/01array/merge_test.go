package array

import (
	"sort"
	"testing"
)

/*
合并两个有序数组

给你两个有序整数数组nums1 和 nums2，请你将 nums2 合并到nums1中，使 nums1 成为一个有序数组。
初始化nums1 和 nums2 的元素数量分别为m 和 n 。你可以假设nums1 的空间大小等于m + n，这样它就有足够的空间保存来自 nums2 的元素。

示例 1：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]

示例 2：
输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]


提示：
nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[i] <= 109

相关标签: 数组, 双指针

*/

// 方法一：直接合并后排序

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 逆向双指针，再原数组上，从后往前赋值
//时间复杂度：O(m+n)。指针移动单调递增，最多移动 m+n 次，因此时间复杂度为 O(m+n)。
//空间复杂度：O(1)。直接对数组 nums1 原地修改，不需要额外空间。
func merge3(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	for m+n > 0 {
		if m == 0 {
			nums1[m+n-1] = nums2[n-1]
			n--
			continue
		}
		if n == 0 {
			nums1[m+n-1] = nums1[m-1]
			m--
			continue
		}

		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

}

// 双指针法，新建一个数组，放最后结果，最后再复制给nums1
//时间复杂度：O(m+n)。指针移动单调递增，最多移动 m+n 次，因此时间复杂度为 O(m+n)。
//空间复杂度：O(m+n)。需要建立长度为 m+n 的中间数组 sorted。
func merge2(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	tmpNums := make([]int, m+n, m+n)
	i, j := 0, 0
	for t := 0; t < m+n; t++ {
		if j >= n {
			tmpNums[t] = nums1[i]
			i++
			continue
		}
		if i >= m {
			tmpNums[t] = nums2[j]
			j++
			continue
		}

		if nums1[i] < nums2[j] {
			tmpNums[t] = nums1[i]
			i++
		} else {
			tmpNums[t] = nums2[j]
			j++
		}
	}

	copy(nums1, tmpNums)

}

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{0, 0, 0}, 0, []int{2, 5, 6}, 3, []int{2, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{1, 2, 4, 5, 6, 0}, 5, []int{3}, 1, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		for i := 0; i < len(tt.expected); i++ {
			if tt.expected[i] != tt.nums1[i] {
				t.Errorf("merge failed. expect: %v, actual: %v", tt.expected, tt.nums1)
				break
			}
		}
	}
}
