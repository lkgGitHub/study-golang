package singleton

import (
	"runtime"
	"sync"
	"testing"
)

func TestGetSingleton(t *testing.T) {
	n := 8
	wg := &sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			GetSingleton()
			runtime.NumGoroutine()
			wg.Done()
		}()
	}
	wg.Wait()
}
