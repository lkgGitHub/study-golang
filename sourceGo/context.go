package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//context.WithCancel(ctx)
	//go handle(ctx, 500*time.Millisecond)
	//select {
	//case <-ctx.Done():
	//	fmt.Println("main", ctx.Err())
	//}
	ctx1 := context.WithValue(context.Background(), "111", "111")

	ctx := context.WithValue(ctx1, "222", "222")
	go func(ctx context.Context) {
		fmt.Println("111: ", ctx.Value("111"))
		fmt.Println("222: ", ctx.Value("222"))
		fmt.Println("000: ", ctx.Value("000"))
	}(ctx)

	time.Sleep(1 * time.Second)

	ctx2, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println("default")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx2)
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
