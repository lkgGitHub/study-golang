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
