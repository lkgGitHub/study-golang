package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// 上面采用了 context.WithTimeout方法，这次使用 context.WithCancel,并让左边的 _ 改成 cancel
	ctx, cancel := context.WithCancel(context.Background())

	// 开启协程
	go workerWithCancel(ctx)

	//主程序等待五秒，防止自动退出
	//当在函数前加 go 意味着这个函数开了一个协程执行函数，主线程会跳过这段
	//当线程结束，不管字协程是否结束，会自动结束，所以我们等待10秒，字协程等待五秒主线程也不会结束。
	time.Sleep(time.Second * 2)

	// 调用 cancel 函数，ctx.Done() 通道自己会赋值，这样子协程就会break了
	cancel()

	fmt.Println("主线程结束")
}

func main2() {
	// context.WithTimeout,返回一个实现context的结构体，有超时返回的功能
	// context.Background()   是一个固定的参数，表示根节点
	// 5*time.Second  表示设置5秒超时，超过两秒子程序会自动退出。
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// 开启协程
	go worker2(ctx)

	// 主程序等待五秒，防止自动退出
	// 当在函数前加 go 意味着这个函数开了一个协程执行函数，主线程会跳过这段
	// 当线程结束，不管字协程是否结束，会自动结束，所以我们等待10秒，字协程等待五秒主线程也不会结束。
	time.Sleep(time.Second * 10)

	fmt.Println("主线程结束")
}

func worker2(ctx context.Context) {
LOOP: // break Loop 跳到这里，并且跳过for{}
	for {
		select {
		// ctx 传入的时候第二个参数设定了5秒 time.Second，当时间到了，ctx.Done(一个通道)，通道中会传入值，
		// 这样 <-ctx.Done() 就能执行力，否则 select就一直执行default{}
		case <-ctx.Done():
			fmt.Println("ctx.Done()通道经过5秒接受到信息，退出")
			break LOOP
		default:
			fmt.Println("执行 default, 并 sleep 1 秒")
			time.Sleep(time.Second)
		}
	}
}

func workerWithCancel(ctx context.Context) {
LOOP: // break Loop 跳到这里，并且跳过for{}
	for {
		select {
		// 当main调用WithCancel的返回参数cancel，cancel调用的时候，ctx.Done(一个通道)，通道中会传入值
		// 这样 <-ctx.Done() 就能执行力，否则 select就一直执行default{}
		case <-ctx.Done():
			fmt.Println("ctx.Done() chan 接受到信息，退出")
			break LOOP
		default:
			fmt.Println("执行 default, 并 sleep 1 秒")
			time.Sleep(time.Second)
		}
	}
}
