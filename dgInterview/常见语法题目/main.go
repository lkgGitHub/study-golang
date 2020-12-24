package main

import (
	"fmt"
	"runtime"
)

/*
打印后
打印中
打印前
panic: 触发异常

*/
func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func main() {
	deferCall()
}

func Demo11() {
	// byte 其实被 alias 到 uint8 上了。
	var i byte
	go func() {
		for i = 0; i <= 255; i++ { // 条件 'i <= 255' 始终为 'true' . uint8 Range: 0 through 255.
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}
