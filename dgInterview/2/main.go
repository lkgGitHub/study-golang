package main

import (
	"strings"
)

func main() {
	println(solution("ABC"))
}

// 字符串的range 后是 UTF8编码的Unicode码点（rune）序列
func solution(str string) bool {
	if len(str) > 3000 {
		return false
	}
	for _, v := range str {
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true
}

func solution2(str string) bool {
	if len(str) > 3000 {
		return false
	}
	for k, v := range str {
		if strings.Index(str, string(v)) != k {
			return false
		}
	}
	return true
}
