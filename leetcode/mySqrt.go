package main
/*
69. x 的平方根
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
 */


func mySqrt(x int) int {
	high := x
	low := 0
	for true {
		mid := (low + high)/2
		if mid * mid < x {
			low = mid + 1
		} else if mid * mid == x {
			return mid
		} else {
			high = mid - 1
		}
		if low > high{
			break
		}
	}
	return high
}
