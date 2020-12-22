package main

import (
	"fmt"
	"runtime"
	"time"
)

// 因为 ch 未初始化，写和读都会阻塞，之后被第一个协程重新赋值，导致写的ch 都阻塞。
// #goroutines 2
func main() {
	ch := make(chan int)
	go func() {
		ch = make(chan int, 1)
		ch <- 1
		println("over func1")
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		println("execute func2")
		<-ch
		println("over func2")
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
