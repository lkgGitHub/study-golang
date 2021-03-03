package main

import "testing"

func TestDemo(t *testing.T) {
	println("bucketCntBits:", bucketCntBits)
	println("bucketCnt:", bucketCnt)
	a := 8 >> 2
	println("bucketCnt:", a)
}
