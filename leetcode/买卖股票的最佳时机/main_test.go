/*	买卖股票的最佳时机
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/
package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []*struct {
		nums   []int
		expect int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		if maxProfit(tt.nums) != tt.expect {
			t.Errorf("nums: %v; want %d, get: %d", tt.nums, tt.expect, maxProfit(tt.nums))
		}
	}
}

// 原始方法
// 复杂度分析
// 时间复杂度：O(n) O(n)，
// 空间复杂度：O(1) O(1)。
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// 原始方法 1
func maxProfit1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0

	for i, j := 0, 1; j < len(prices); {
		if prices[j] > prices[i] {
			profit = profit + prices[j] - prices[i]
			i++
			j++
		} else {
			i++
			j++
		}
	}
	return profit
}
