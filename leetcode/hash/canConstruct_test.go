package main

import "testing"

/*
383. 赎金信
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：
输入：ransomNote = "a", magazine = "b"
输出：false

示例 2：
输入：ransomNote = "aa", magazine = "ab"
输出：false

示例 3：
输入：ransomNote = "aa", magazine = "aab"
输出：true

提示：
1 <= ransomNote.length, magazine.length <= 105
ransomNote 和 magazine 由小写英文字母组成
*/
func Test_canConstruct(t *testing.T) {
	println('a')

	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"示例 1", args{"a", "b"}, false},
		{"示例 2", args{"aa", "ab"}, false},
		{"示例 3", args{"aa", "aab"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct2(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[int32]int, len(ransomNote))
	for _, v := range magazine {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, r := range ransomNote {
		if _, ok := m[r]; !ok {
			return false
		} else {
			if m[r] == 0 {
				return false
			}
			m[r]--
		}
	}

	return true
}

// 官方结题
func canConstruct2(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-97]++ // ch - 'a'  == ch - 97
	}
	for _, ch := range ransomNote {
		if cnt[ch-97] == 0 {
			return false
		}
		cnt[ch-97]--
	}
	return true
}
