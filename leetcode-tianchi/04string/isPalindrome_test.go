package main

import (
	"strings"
	"testing"
)

/*
	125. 验证回文串 https://leetcode.cn/problems/valid-palindrome/
*/
func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"示例1", args{"A man, a plan, a canal: Panama"}, true},
		{"示例2", args{"race a car"}, false},
		{"示例3", args{"0P"}, false},
		{"示例4", args{" P"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; {
		left := s[i : i+1]
		right := s[j : j+1]

		// 只考虑字母和数字字符，ASCII码 A-Z(65~90) a-z(97~122)
		println(left, right)
		if !((s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122)) {
			i++
			continue
		}
		if !((s[j] >= 48 && s[j] <= 57) || (s[j] >= 65 && s[j] <= 90) || (s[j] >= 97 && s[j] <= 122)) {
			j--
			continue
		}

		if strings.ToLower(left) == strings.ToLower(right) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
