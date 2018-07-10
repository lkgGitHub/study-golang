package main

import "strings"
import "fmt"

func Index(vs []string, t string) int {
	for i, v := range vs{
		if v == t {
			return i
		}
	}
}

func Include(vs []strings, t strings.Reader)  {
	
}