package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

/*
上面这道题在不执行resp.Body.Close()的情况下，泄漏了吗？如果泄漏，泄漏了多少个goroutine?
答案：虽然执行了 6 次循环，而且每次都没有执行 Body.Close() ,就是因为执行了ioutil.ReadAll()把内容都读出来了，连接得以复用，
	 因此只泄漏了一个读goroutine和一个写goroutine，最后加上main goroutine，所以答案就是3个goroutine。
*/
func main() {
	num := 6
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.baidu.com")
		_, _ = ioutil.ReadAll(resp.Body)
	}
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}
