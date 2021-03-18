package string

import (
	"fmt"
	"testing"
)

/*
反转字符串

编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

示例 1：
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]

示例 2：
输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]

相关标签: 双指针 字符串
*/

// 双指针
// 时间复杂度：O(N)，其中 NN 为字符数组的长度。一共执行了 N/2 次的交换。
// 空间复杂度：O(1)。只使用了常数空间来存放若干变量。
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

func reverseString2(s []byte) {
	var tmp byte
	var i int
	for j := len(s) - 1; i < j; {
		tmp = s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		s        []byte
		expected []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		reverseString(tt.s)
		isPass := true
		for i := 0; i < len(tt.expected); i++ {
			if tt.expected[i] != tt.s[i] {
				isPass = false
			}
		}
		if !isPass {
			t.Error(fmt.Sprintf("TestRotate failed. expect: %v , actual: %v", tt.expected, tt.s))
		}
	}
}
