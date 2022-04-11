package hash

import "testing"

/*
161. 存在重复元素 II
难度：简单
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [1,2,3,1], k = 3
输出：true

示例 2：
输入：nums = [1,0,1,1], k = 1
输出：true

示例 3：
输入：nums = [1,2,3,1,2,3], k = 2
输出：false




提示：
1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
*/
func TestContainsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "示例2",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "示例3",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("ContainsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 方法一：哈希表，原始版本
func containsNearbyDuplicate1(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := m[n]; ok {
			if i-j <= k {
				return true
			}
			m[n] = i // 遇到重复数字，需要更新下标
		} else {
			m[n] = i
		}
	}
	return false
}

// 方法一：哈希表，精简版
func containsNearbyDuplicate2(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := m[n]; ok && i-j <= k {
			return true
		}
		m[n] = i
	}
	return false
}

// todo 方法二：滑动窗口
