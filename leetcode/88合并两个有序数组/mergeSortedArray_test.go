package mergeSortedArray

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	println("start")
	var tests = []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		{[]int{1}, 1, []int{}, 0},
		{[]int{2, 0}, 1, []int{1}, 1},
	}
	for i, test := range tests{
		println("i:", i)
		merge(test.nums1, test.m, test.nums2, test.n)
	}
}

// 双指针法，利用短路|| 来避免切片指针越界
func merge(nums1 []int, m int, nums2 []int, n int)  {
	var result []int
	i := 0
	j := 0
	for ; ; {
		if i == m && j == n {
			copy(nums1, result)
			break
		}
		if j == n || (i < m  && nums1[i] <= nums2[j] ){
			result = append(result, nums1[i])
			i++
		}else if i == m || (j < n && nums1[i] > nums2[j])  {
			result = append(result, nums2[j])
			j++
		}
	}
	fmt.Printf("===>%+v \n", nums1)

}
