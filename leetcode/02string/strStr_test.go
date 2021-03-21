package string

import (
	"fmt"
	"strings"
	"testing"
)

/*
实现 strStr()
// todo: Rabin Karp - 常数复杂度,暂时放弃
// todo: KMP 算法详解
实现strStr()函数。
给定一个haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当needle是空字符串时我们应当返回 0 。这与C语言的strstr()以及 Java的indexOf()定义相符。

相关标签: 双指针	字符串
*/

// 方法一：子串逐一比较
func strStr1(haystack string, needle string) int {
	h, n := len(haystack), len(needle)
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

// 方法二：双指针
func strStr2(haystack string, needle string) int {
	h, n := len(haystack), len(needle)
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

// lkg
func MyStrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	j := 0    // needle 指针位置
	tmp := -1 // 记录第一个字符相等的位置
	for i := 0; i < len(haystack); i++ {
		if j == len(needle) {
			return tmp
		}
		if haystack[i] == needle[j] {
			if j == 0 {
				tmp = i
			}
			j++
		} else if tmp != -1 {
			// 注意整体回退！
			tmp, i, j = -1, tmp, 0
		}
	}
	if j == len(needle) {
		return tmp
	}
	return -1
}

// go 官方包
func strGolang(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		//{"hello", "ll", 2},
		//{"aaaaa", "bba", -1},
		//{"leetcode", "", 0},
		{"aabaabcd", "abc", 4},
		{"abc", "bc", 1},
		{"abc", "bcd", -1},
		{"mississippi", "issip", 4},
	}
	for _, tt := range tests {
		actual := MyStrStr(tt.haystack, tt.needle)
		if tt.expected != actual {
			t.Error(fmt.Sprintf("strStr failed. expect: %v , actual: %v", tt.expected, actual))
		}
	}
}
