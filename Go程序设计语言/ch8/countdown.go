package main

import (
	"fmt"
	"os"
	"time"
)

// 模拟火箭发射倒计时，控制台输入回车后，停止倒计时
func Launch() {
	abort := make(chan struct{})
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second) // “time.Tick函数返回一个channel”
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			fmt.Println("Do nothing")
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	//fmt.Println("Commencing countdown.  Press return to abort.")
	//select {
	//case <-time.After(10*time.Second): // time.After函数会立即返回一个channel，并起一个新的goroutine在经过特定的“时间后向该channel发送一个独立的值
	//	fmt.Println("do nothing!")
	//case <-abort :
	//	fmt.Println("Launch aborted!")
	//	return
	//}
}

// ch这个channel的buffer大小是1，所以会交替的为空或为满，所以只有一个case可以进行下去，无论i是奇数或者偶数，它都会打印0 2 4 6 8
// 增加前一个例子的buffer大小会使其输出变得不确定，因为当buffer既不为满也不为空时，select语句的执行情况就像是抛硬币的行为一样是随机的。
func testSelect() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("x: ", x)
		case ch <- i:
			println("i: ", i)
		}
	}
}


