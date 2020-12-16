package main

import (
	"strings"
	"unicode"
)

func main() {
	solution("a b c")
}

// unicode.IsLetter判断字符是否是字母，之后使用strings.Replace来替换空格
func solution(s string) (string, bool) {
	strArr := []rune(s)
	if len(strArr) > 1000 {
		return s, false
	}
	for _, v := range strArr {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	return strings.ReplaceAll(s, " ", "%20"), true
}
