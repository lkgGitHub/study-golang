package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

func TestFixedWindowRateLimiting(t *testing.T) {
	fixedWindowRateLimiting()
}

func TestTokenBucketRateLimiting(t *testing.T) {
	tokenBucketRateLimiting()
}

// 固定窗口 fixedWindowRateLimiting
func fixedWindowRateLimiting() {
	// 速率限制为每秒100个请求
	limiter := rate.NewLimiter(rate.Limit(100), 1) // 允许每秒100次

	for i := 0; i < 200; i++ {
		// limiter.Allow()，如果允许请求，则返回true;如果超过速率限制，则返回false
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		go process()
	}
}

func tokenBucketRateLimiting() {
	// 其速率限制为每秒 10 个请求，突发 5 个请求
	// 代表每秒可以向 Token 桶中产生多少 token。b 代表 Token 桶的容量大小
	limiter := rate.NewLimiter(rate.Limit(10), 5)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
	for i := 0; i < 200; i++ {
		// limiter.Wait() 方法，该方法会阻塞直到有令牌可用
		if err := limiter.Wait(ctx); err != nil {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		go process()
	}
}

// 动态速率限制
func dynamicRateLimiting() {
	limiter := rate.NewLimiter(rate.Limit(10), 1) // 允许每秒100次

	// Dynamic rate adjustment
	go func() {
		time.Sleep(time.Second * 10) // 每10秒调整 limiter
		fmt.Println("---adjust limiter---")
		limiter.SetLimit(rate.Limit(200)) // 将 limiter 提升到每秒 200
	}()

	for i := 0; i < 3000; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		process()
	}
}

// 自适应速率限制
func adaptiveRateLimiting() {
	limiter := rate.NewLimiter(rate.Limit(10), 1) // 允许每秒10次

	// 自适应调整
	go func() {
		for {
			time.Sleep(time.Second * 10)
			responseTime := measureResponseTime() // 测量之前请求的响应时间
			if responseTime > 500*time.Millisecond {
				fmt.Println("---adjust limiter 50---")
				limiter.SetLimit(rate.Limit(50))
			} else {
				fmt.Println("---adjust limiter 100---")
				limiter.SetLimit(rate.Limit(100))
			}
		}
	}()

	for i := 0; i < 3000; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		process()
	}
}

// 测量以前请求的响应时间
// 执行自己的逻辑来测量响应时间
func measureResponseTime() time.Duration {
	return time.Millisecond * 100
}

// 处理请求
func process() {
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond) // 模拟请求处理时间
}
