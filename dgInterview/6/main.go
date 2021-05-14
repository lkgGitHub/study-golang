package main

import (
	"math/rand"
	"sync"
	"time"
)

/*
写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
*/

// goroutine 在golang中式非阻塞的
// channel 无缓冲情况下，读写都是阻塞的，且可以用for循环来读取数据，当管道关闭后，for 退出。
// golang 中有专用的select case 语法从管道读取数据。
func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int) {
		defer wg.Done()
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(10)
		}
		close(ch)
	}(ch)

	go func(ch chan int) {
		defer wg.Done()
		for i := range ch {
			println(i)
		}
		//for {
		//	i, open := <-ch
		//	if open {
		//		println(i)
		//	} else {
		//		return
		//	}
		//}
	}(ch)
	wg.Wait()
}
