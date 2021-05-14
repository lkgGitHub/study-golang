package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		// 要求每秒钟执行一次，首先想到的就是 time.Ticker 对象，该函数可每秒钟往chan中放一个Time
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <- ticker.C:
				go func() {
					defer func() {
						// 捕获 panic 一般会用到 recover() 函数
						if p := recover(); p!=nil{
							fmt.Printf("internal error: %v \n", p)
						}
					}()
					proc()
				}()
			}
		}
	}()
	wg.Wait()
}

func proc() {
	panic("ok")
}
