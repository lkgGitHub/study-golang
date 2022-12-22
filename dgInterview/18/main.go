package main

import (
	"fmt"
	"sync"
	"time"
)

/*
   算法:为 sync.WaitGroup 中Wait函数支持 WaitTimeout 功能
   eg:假设⼀个API接⼝，接收⼀个图⽚URL的列表，然后批量下载图⽚，接⼝的超时时间是5s，完成下⾯的代码⽚段。
*/
func main() {
	// 批量下载图片
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			// 假设下载图片的时间
			fmt.Println("Sleep: ", num, " second")
			time.Sleep(time.Duration(num) * time.Second)
		}(i)
	}
	// 设置超时时间
	if isTimeout(&wg, time.Second*5) {
		fmt.Println("Timeout")
		return
	}
	fmt.Println("Done")
	time.Sleep(time.Second * 10)
}
func isTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	return false
	// 你的代码
}
