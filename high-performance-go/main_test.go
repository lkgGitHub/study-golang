package main

import (
	"sort"
	"testing"
)

func main() {

}
func TestCom(t *testing.T) {
	//s1 := []int{3,2,4,1}
	//s2 := []int{1,3,5,2}
	//fmt.Println(combineList(s1, s2))
	println(com("marginle", "valaienie"))
}

func com(str1, str2 string) string {
	var str []rune
	m := make(map[rune]struct{}, len(str2))
	for _, s := range str2{
		m[s] = struct{}{}
	}

	strMap := make(map[rune]struct{}, len(str1))
	for _, s := range str1{
		if _, ok := m[s]; ok {
			if _, has := strMap[s];!has {
				strMap[s] =struct{}{}
				str = append(str, s)
			}
		}
	}
	return  string(str)
}

func combineList(s1, s2 []int) []int  {
	s := make([]int, 0, len(s1) + len(s2))

	for i := 0; i < len(s1); i++ {
		s = append(s, s1[i])
	}
	for i := 0; i < len(s2); i++ {
		s = append(s, s2[i])
	}

	sort.Ints(s)
	return s
}