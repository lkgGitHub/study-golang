package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/power-of-two/solution/2de-mi-by-leetcode-solution-rny3/
/*
171. 2 的幂
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。

示例 1：

输入：n = 1
输出：true
解释：20 = 1
示例 2：

输入：n = 16
输出：true
解释：24 = 16
示例 3：

输入：n = 3
输出：false
示例 4：

输入：n = 4
输出：true
示例 5：

输入：n = 5
输出：false

提示：
-231 <= n <= 231 - 1
进阶：你能够不使用循环/递归解决此问题吗？
*/
// 如果是2的幂，二进制只有一个1，所以去掉一个1与其相与为0  &
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// 一个数 n 是 2 的幂，当且仅当 n 是正整数，并且 n 的二进制表示中仅包含 1 个 1。
func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}

	s := fmt.Sprintf("%b", n)
	return strings.Count(s, "1") == 1
}

// 暴力方法，内存占用过高
func isPowerOfTwo3(n int) bool {
	if n <= 0 {
		return false
	}

	if n%2 != 0 && n != 1 {
		return false
	}
	n = n / 2
	if n == 1 {
		return true
	}
	return true
}
