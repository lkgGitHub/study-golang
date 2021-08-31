package main

import (
	"fmt"
	"strings"
	"testing"
)

// go test -bench="Concat$" -benchmem .
// -bench 参数支持传入一个正则表达式，匹配到的用例才会得到执行，例如，只运行以 Concat 结尾的
// -benchmem 参数可以度量内存分配的次数。内存分配次数也性能也是息息相关的，例如不合理的切片容量，将导致内存重新分配，带来不必要的开销。

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

func BenchmarkPlusConcat(b *testing.B)       { benchmark(b, plusConcat) }
func BenchmarkSprintfConcat(b *testing.B)    { benchmark(b, sprintfConcat) }
func BenchmarkBuilderConcat(b *testing.B)    { benchmark(b, builderConcat) }
func BenchmarkPreBuilderConcat(b *testing.B) { benchmark(b, preBuilderConcat) }
func BenchmarkBufferConcat(b *testing.B)     { benchmark(b, bufferConcat) }
func BenchmarkByteConcat(b *testing.B)       { benchmark(b, byteConcat) }
func BenchmarkPreByteConcat(b *testing.B)    { benchmark(b, preByteConcat) }

/*
goos: windows
goarch: amd64
pkg: study-goland/high-performance-go/02string
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkPlusConcat-16                        22          51450727 ns/op        530997428 B/op     10018 allocs/op
BenchmarkSprintfConcat-16                     12          90441042 ns/op        834690504 B/op     37472 allocs/op
BenchmarkBuilderConcat-16                  13131             90776 ns/op          505840 B/op         24 allocs/op
BenchmarkPreBuilderConcat-16               26695             44837 ns/op          106496 B/op          1 allocs/op
BenchmarkBufferConcat-16                   13053             92072 ns/op          423536 B/op         13 allocs/op
BenchmarkByteConcat-16                     12832             93496 ns/op          612337 B/op         25 allocs/op
BenchmarkPreByteConcat-16                  25221             47897 ns/op          212994 B/op          2 allocs/op
PASS
ok      study-goland/high-performance-go/02string       12.331s
*/

// builder.Cap() 查看字符串拼接过程中，strings.Builder 的内存申请过程。
// 16 32 64 128 256 512 1024 1280 1792 2304 3072 4096 5376 6784 9472 12288 16384 20480 27264 40960 57344 73728 98304 122880 --- PASS: TestBuilderConcat (0.00s)
func TestBuilderConcat(t *testing.T) {
	var str = randomString(10)
	var builder strings.Builder
	builderCap := 0
	for i := 0; i < 10000; i++ {
		if builder.Cap() != builderCap {
			fmt.Print(builder.Cap(), " ")
			builderCap = builder.Cap()
		}
		builder.WriteString(str)
	}
}
