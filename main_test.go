package study_golang

import (
	"fmt"
	"testing"
)

func TestByteToInt(t *testing.T) {
	fmt.Println("Hello, ", Bytes2Bits([]byte{2, 1}))
}

func Bytes2Bits(data []byte) []uint8 {
	dst := make([]uint8, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint8(7 - i)
			dst = append(dst, (v>>move)&1)
		}
	}
	fmt.Println(len(dst))
	return dst
}

func TestStringer(t *testing.T) {
	//var s fmt.Stringer
	//s = "s"
	//fmt.Println(s)
	// 编译错误

	const X = 7.0
	var x interface{} = X
	if y, ok := x.(int); ok {
		fmt.Println(y)
	} else {
		fmt.Println(int(y))
	}
	// 最后输出的是 0，因为 y 类型断言失败，则 y 为 int 的零值
}
