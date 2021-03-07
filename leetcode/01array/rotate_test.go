package array

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode-solution-nipk/
// 旋转数组

// 方法1： 每次右移一位
func rotate(nums []int, k int) {
	k = k % len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
	fmt.Println("nums:", nums)
}

//方法1： 直接移动k位
func rotate2(nums []int, k int) {
	k = k % len(nums)

	numsTemp := make([]int, 0, k)
	for i, n := range nums {
		if i < len(nums)-k {
			numsTemp = append(numsTemp, n)
		} else {
			nums[i+k-len(nums)] = n
		}
	}

	for i, n := range numsTemp {
		nums[k+i] = n
	}

	fmt.Println("nums:", nums)
}

func TestRotate(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{-1}, 2, []int{-1}},
	}
	for _, tt := range tests {
		rotate2(tt.nums, tt.k)
		if !equalInt(tt.expect, tt.nums) {
			t.Error(fmt.Sprintf("TestRotate failed. expect: %v , actual: %v", tt.expect, tt.nums))
		}
	}
}

func equalInt(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestDemo(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst)

	//// 设置元素数量为1000
	//const elementCount = 10
	//// 预分配足够多的元素切片
	//srcData := make([]int, elementCount)
	//// 将切片赋值
	//for i := 0; i < elementCount; i++ {
	//	srcData[i] = i
	//}
	//// 引用切片数据
	//refData := srcData
	//// 预分配足够多的元素切片
	//copyData := make([]int, elementCount)
	//// 将数据复制到新的切片空间中
	//copy(copyData, srcData)
	//// 修改原始数据的第一个元素
	//srcData[0] = 9
	//// 打印引用切片的第一个元素
	//fmt.Println(refData[0])
	//// 打印复制切片的第一个和最后一个元素
	//fmt.Println(copyData[0], copyData[elementCount-1])
	//// 复制原始数据从4到6(不包含)
	//copy(copyData, srcData[4:6])
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("%d ", copyData[i])
	//}
}
