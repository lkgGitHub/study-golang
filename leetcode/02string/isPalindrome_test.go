package string

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

/*
验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true

示例 2:
输入: "race a car"
输出: false
相关标签: 双指针	字符串
*/
func isPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i <= j; i++ {
		for i <= j {
			if x := s[i]; (x >= '0' && x <= '9') || (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') {
				break
			} else {
				i++
			}
		}
		for i <= j {
			if x := s[j]; (x >= '0' && x <= '9') || (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') {
				break
			} else {
				j--
			}
		}
		if i > j {
			break
		}
		if strings.ToLower(string(s[i])) == strings.ToLower(string(s[j])) {
			//fmt.Printf("s[i]: %d \n", s[i])
			//fmt.Printf("s[j]: %d \n", s[j])
			//println(s[i] == s[j])
			//println(s[i]-s[j] == 32)
			//println(s[j]-s[i] == 32)
			j--
		} else {
			//fmt.Printf("s[i]: %s \n", string(s[i]))
			fmt.Printf("s[j]: %s \n", string(s[j]))
			//fmt.Printf("s[i]: %d \n", s[j])
			fmt.Printf("s[j]: %d \n", s[j])
			return false
		}

	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		//{"A man, a plan, a canal: Panama", true},
		//{"race a car", false},
		//{"", true},
		//{"A man, a plan, a canal: Panama", true},
		//{".,", true},
		//{"0P", false},
		{"ab_a", true},
	}
	for _, tt := range tests {
		actual := isPalindrome(tt.s)
		if tt.expected != actual {
			t.Error(fmt.Sprintf("TestRotate failed. expect: %v , actual: %v; s: %s",
				tt.expected, actual, tt.s))
		}
	}
}
func TestIsPalindromeDemo(t *testing.T) {
	a := 'z'
	A := 'A'
	fmt.Printf("0: %d \n", '0')
	fmt.Printf("9: %d \n", '9')
	fmt.Printf("a: %d \n", 'a')
	fmt.Printf("z: %d \n", 'z')
	fmt.Printf("A: %d \n", 'A')
	fmt.Printf("Z: %d \n", 'Z')
	println(a)
	println(A)
	println(a == A)
	fmt.Printf("a - A: %d \n", 'a'-'A')
	fmt.Printf("A - a: %d \n", 'A'-'a')
	fmt.Printf("A - a: %v \n", 'A'-'a' == -32)
	fmt.Printf("0 - p: %v \n", '0'-'p')
	println(math.Abs(float64('a'-'A')) == 32)
	println(int(math.Abs(float64(65 - 97))))

}
