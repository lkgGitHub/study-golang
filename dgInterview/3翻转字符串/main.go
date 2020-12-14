package main

import (
	"fmt"
	"strings"
)

/*
文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列
因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的：
s[0] = 'L' // compile error: cannot assign to s[0]
*/
func main() {
	s := "abcde"
	if r, ok := solution2(s); ok {
		println(r)
	} else {
		println("error")
	}

}

func solution(s string) (string, bool) {
	str := []rune(s)
	size := len(str)
	if size > 5000 {
		return "", false
	}
	for i := 0; i < size/2; i++ {
		str[i], str[size-i-1] = str[size-i-1], str[i]
	}
	fmt.Printf("%v \n", str)
	return string(str), true
}

func solution2(s string) (string, bool) {
	strArr := strings.Split(s, "")
	size := len(strArr)
	if size > 5000 {
		return "", false
	}
	for i := 0; i < size/2; i++ {
		strArr[i], strArr[size-i-1] = strArr[size-i-1], strArr[i]
	}
	fmt.Printf("%v \n", strArr)
	return strings.Join(strArr, ""), true
}

func reverseString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str), true
}
