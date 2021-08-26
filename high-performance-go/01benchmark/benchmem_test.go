package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	numbs := make([]int, 0, n)
	for i:=0;i<n;i++{
		numbs = append(numbs, rand.Int())
	}
	return numbs
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	numbs := make([]int, 0)
	for i := 0; i < n; i++ {
		numbs = append(numbs, rand.Int())
	}
	return numbs
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(1000000)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(1000000)
	}
}

/*
goos: windows
goarch: amd64
pkg: study-goland/high-performance-go/01benchmark
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkGenerateWithCap-16           75          14391067 ns/op         8003681 B/op          1 allocs/op
BenchmarkGenerate-16                  62          17999784 ns/op        45188659 B/op         42 allocs/op
PASS
ok      study-goland/high-performance-go/01benchmark    2.389s


allocs/op  内存分配次数。

 */