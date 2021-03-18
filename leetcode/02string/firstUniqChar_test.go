package string

import (
	"fmt"
	"testing"
)

func firstUniqChar(s string) int {
	m := make(map[rune]int, len(s))
	for _, r := range s {
		if v, ok := m[r]; ok {
			m[r] = v + 1
		} else {
			m[r] = 1
		}
	}

	for i, r := range s {
		if m[r] == 1 {
			return i
		}
	}
	return -1
}

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabbcc", -1},
	}
	for _, tt := range tests {
		actual := firstUniqChar(tt.s)
		if tt.expected != actual {
			t.Error(fmt.Sprintf("TestRotate failed. expect: %v , actual: %v", tt.expected, actual))
		}
	}
}
