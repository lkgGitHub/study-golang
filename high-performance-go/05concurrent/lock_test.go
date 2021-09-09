package concurrent

import (
	"sync"
	"testing"
)

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkReadMore(b *testing.B)    { benchmark(b, &Lock{}, 9, 1) }
func BenchmarkReadMoreRW(b *testing.B)  { benchmark(b, &RWLock{}, 9, 1) }
func BenchmarkWriteMore(b *testing.B)   { benchmark(b, &Lock{}, 1, 9) }
func BenchmarkWriteMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 1, 9) }
func BenchmarkEqual(b *testing.B)       { benchmark(b, &Lock{}, 5, 5) }
func BenchmarkEqualRW(b *testing.B)     { benchmark(b, &RWLock{}, 5, 5) }


/*
每个读写操作耗时 1 微秒(百万分之一秒)结论如下：
读写比为 9:1 时，读写锁的性能约为互斥锁的 8 倍
读写比为 1:9 时，读写锁性能相当
读写比为 5:5 时，读写锁的性能约为互斥锁的 2 倍

goos: darwin
goarch: amd64
pkg: study-goland/high-performance-go/05concurrent
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkReadMore-16       	     250	   4910003 ns/op
BenchmarkReadMoreRW-16     	    1644	    738845 ns/op
BenchmarkWriteMore-16      	     247	   4795608 ns/op
BenchmarkWriteMoreRW-16    	     277	   4307996 ns/op
BenchmarkEqual-16          	     248	   4787261 ns/op
BenchmarkEqualRW-16        	     481	   2508251 ns/op
PASS
ok  	study-goland/high-performance-go/05concurrent	9.445s
 */