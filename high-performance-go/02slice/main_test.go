package main

import (
	"runtime"
	"testing"
)

/*
可能存在只使用了一小段切片，但是底层数组仍被占用，得不到使用，推荐使用 copy 替代默认的 re-slice
*/

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024) // 1M。 128*1024 个 int 整型，恰好为 1 MB
		ans = append(ans, f(origin))
		runtime.GC()
	}
	printMem(t)
	_ = ans
}

func TestLastCharsBySlice(t *testing.T) { testLastChars(t, lastNumsBySlice) }
func TestLastCharsByCopy(t *testing.T)  { testLastChars(t, lastNumsByCopy) }

/*
=== RUN   TestLastCharsBySlice
    main_test.go:16: 100.25 MB
--- PASS: TestLastCharsBySlice (0.20s)
=== RUN   TestLastCharsByCopy
    main_test.go:17: 2.26 MB
--- PASS: TestLastCharsByCopy (0.19s)
PASS
ok      study-goland/high-performance-go/02slice        0.636s
*/
