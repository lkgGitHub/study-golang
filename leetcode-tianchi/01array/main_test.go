package array

import (
	"fmt"
	"sort"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"示例1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}},
		{"示例2", args{[]int{}}, [][]int{}},
		{"示例3", args{[]int{0}}, [][]int{}},
		{"示例4", args{[]int{0, 0, 0, 0}}, [][]int{
			{0, 0, 0},
		}},
		{"示例5", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}},
		{"", args{[]int{1, 2, -2, -1}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum2(tt.args.nums)
			m := make(map[string]struct{}, len(tt.want))
			for _, w := range tt.want {
				m[fmt.Sprintf("%d%d%d", w[0], w[1], w[2])] = struct{}{}
			}
			for _, g := range got {
				if len(g) != 3 {
					t.Errorf("threeSum() = %v, want %v", got, tt.want)
					return
				}

				if _, ok := m[fmt.Sprintf("%d%d%d", g[0], g[1], g[2])]; ok {
					delete(m, fmt.Sprintf("%d%d%d", g[0], g[1], g[2]))
				} else {
					t.Errorf("threeSum() = %v, want %v", got, tt.want)
				}
			}
			if len(m) > 0 {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	var results [][]int
	nums1 := make([]int, 0, len(nums)/2)
	nums2 := make([]int, 0, len(nums)/2)
	nums0 := make([]int, 0, 3)
	for _, n := range nums {
		switch true {
		case n < 0:
			nums1 = append(nums1, n)
		case n == 0:
			if len(nums0) >= 3 {
				continue
			}
			nums0 = append(nums0, n)
		case n > 0:
			nums2 = append(nums2, n)
		}
	}
	if len(nums0) >= 3 {
		results = append(results, []int{0, 0, 0})
	}
	if len(nums1) == 0 || len(nums2) == 0 {
		return results
	}

	for i := 0; i < len(nums); i++ {
		m := make(map[int]int, len(nums)-1)
		target := nums[i]

		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			m[nums[j]] = j
		}

		for j := i; j < len(nums); j++ {
			if i == j {
				continue
			}
			if k, ok := m[0-target-nums[j]]; ok && k != j {
				results = append(results, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	m := make(map[string]struct{}, len(results))
	for i := 0; i < len(results); {
		r := results[i]
		sort.Ints(r)
		k := fmt.Sprintf("%d%d%d", r[0], r[1], r[2])
		if _, ok := m[k]; ok {
			results = append(results[0:i], results[i+1:]...)
		} else {
			m[k] = struct{}{}
			i++
		}
	}

	return results
}

func threeSum1(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	var results [][]int
	for i := 0; i < len(nums); i++ {
		m := make(map[int]int, len(nums)-1)
		target := nums[i]

		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			m[nums[j]] = j
		}

		for j := i; j < len(nums); j++ {
			if i == j {
				continue
			}
			if k, ok := m[0-target-nums[j]]; ok && k != j {
				results = append(results, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	m := make(map[string]struct{}, len(results))
	for i := 0; i < len(results); {
		r := results[i]
		sort.Ints(r)
		k := fmt.Sprintf("%d%d%d", r[0], r[1], r[2])
		if _, ok := m[k]; ok {
			results = append(results[0:i], results[i+1:]...)
		} else {
			m[k] = struct{}{}
			i++
		}
	}

	return results
}
