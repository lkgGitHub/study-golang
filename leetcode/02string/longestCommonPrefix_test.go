package string

import (
	"fmt"
	"testing"
)

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：
0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	size := len(strs[0])
	for _, str := range strs {
		if size > len(str) {
			size = len(str)
		}
	}
	i := 0
loop:
	for ; i < size; i++ {
		s := strs[0][i]
		for _, str := range strs {
			if s != str[i] {
				break loop
			}
		}
	}
	return strs[0][:i]

}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		//{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"dog", "", "car"}, ""},
	}
	for _, tt := range tests {
		actual := longestCommonPrefix(tt.strs)
		if tt.expected != actual {
			t.Error(fmt.Sprintf("strStr failed. expect: %v , actual: %v", tt.expected, actual))
		}
	}
}

func TestTestLongestCommonPrefixDemo(t *testing.T) {
	println("aaaa"[:2])
}
