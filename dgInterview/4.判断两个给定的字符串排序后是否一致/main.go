package main

import "strings"

func main() {
	s1 := "abc"
	s2 := "abcd"
	println(solution(s1, s2))
}

func solution(s1, s2 string) bool {
	if len(s1) > 5000 || len(s2) > 5000 || len(s1) != len(s2) {
		return false
	}

	for _, r := range s1 {
		if strings.Count(s2, string(r)) <= 0 {
			return false
		}
	}

	return true
}
