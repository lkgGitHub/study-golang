package string

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

/*
整数反转

给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31− 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：
输入：x = 123
输出：321

示例 2：
输入：x = -123
输出：-321

示例 3：
输入：x = 120
输出：21

示例 4：
输入：x = 0
输出：0

提示：
-231 <= x <= 231 - 1

相关标签:	数学
*/
// todo: 弹出和推入数字 & 溢出前进行检查

// 转换成字符串，利用字符串反转。最后再转成int32，超过int32的转换过程中会有error，直接返回0
func reverse(x int) int {
	xStr := strconv.Itoa(x)
	yStr := strings.Builder{}
	if xStr[0] == '-' {
		yStr.WriteRune('-')
		xStr = xStr[1:]
	}
	for i := len(xStr) - 1; i >= 0; i-- {
		yStr.WriteByte(xStr[i])
	}
	y, err := strconv.ParseInt(yStr.String(), 10, 32)
	if err != nil {
		return 0
	}
	return int(y)
}
func TestReverse(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
	}
	for _, tt := range tests {
		actual := reverse(tt.x)
		if tt.expected != actual {
			t.Error(fmt.Sprintf("reverse failed. expect: %v , actual: %v", tt.expected, actual))
		}
	}
}

func TestReverseDemo2(t *testing.T) {
	println("12345 / 10000", 12345/10000)
	println("12345 / 1000", 12345/1000)
	println("12345 / 100", 12345/100)
	println("12345 / 10", 12345/10)

	println("12345 % 10000", 12345%10000)
	println("12345 % 1000", 12345%1000)
	println("12345 % 100", 12345%100)
	println("12345 % 10", 12345%10)
}

func TestReverseDemo(t *testing.T) {
	maxInt64 := math.MaxInt32
	i, err := strconv.ParseInt(strconv.Itoa(maxInt64+1), 10, 32)
	if err != nil {
		t.Error(err.Error())
		return
	}
	println("===>", i)
}
