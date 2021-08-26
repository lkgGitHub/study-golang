package main

import "testing"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

/*
goos: windows
goarch: amd64
pkg: study-goland/high-performance-go/01benchmark
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkFib
BenchmarkFib-16    	     280	   4242482 ns/op
PASS

BenchmarkFib-16  16核心cpu
可以-cpu 参数改变 GOMAXPROCS

执行了280次，


-benchtime	benchmark 的默认时间是 1s，那么我们可以使用 -benchtime 指定为 5s。
-count 		参数可以用来设置 benchmark 的轮数。例如， -count=3 进行 3 轮 benchmark。


*/