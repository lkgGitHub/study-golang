package string

import (
	"fmt"
	"sort"
	"testing"
)

/*
有效的字母异位词

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

相关标签: 排序	哈希表
*/

// 排序
func isAnagram(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// 哈希表2, 利用只有26个字幕，可以用数组的位来代替字母
// 从另一个角度考虑，tt 是 ss 的异位词等价于「两个字符串中字符出现的种类和次数均相等」。
// 由于字符串只包含 26 个小写字母，因此我们可以维护一个长度为 26 的频次数组 table，先遍历记录字符串 s 中字符出现的频次，
// 然后遍历字符串 t，减去 table 中对应的频次，如果出现 table[i]<0，则说明 t 包含一个不在 s 中的额外字符，返回 false 即可。
func isAnagram1(s string, t string) bool {
	var c1, c2 [26]uint8
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

// 哈希表1，可处理进阶题目
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]uint16, len(s))
	for _, i := range s {
		if _, ok := sMap[i]; ok {
			sMap[i] += 1
		} else {
			sMap[i] = 1
		}
	}

	for _, i := range t {
		if _, ok := sMap[i]; ok {
			if sMap[i] == 0 {
				return false
			} else {
				sMap[i] -= 1
			}
		} else {
			return false
		}
	}

	return true
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}
	for _, tt := range tests {
		actual := isAnagram(tt.s, tt.t)
		if tt.expected != actual {
			t.Error(fmt.Sprintf("TestRotate failed. expect: %v , actual: %v", tt.expected, actual))
		}
	}
}

func TestIsAnagramDemo(t *testing.T) {
	s := "aA" // \"nagaram\"\n"
	println(s[0])
	println(s[1])
	println(' ')
	println(len(s))
}
