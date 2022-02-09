package main

import (
	"sync"
)

// 全局变量
var counter int
var lock sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			counter++
			lock.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}
