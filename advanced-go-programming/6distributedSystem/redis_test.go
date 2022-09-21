package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	lockKey    = "counter_lock"
	counterKey = "counter"
)

func TestRedisSetNX(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
}

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.2.1242:6378",
	})
	var ctx = context.Background()

	// lock
	resp := client.SetNX(ctx, lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(ctx, counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp = client.SetNX(ctx, counterKey, cntValue, 0)
		_, err = resp.Result()
		if err != nil {
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)
	delResp := client.Del(ctx, lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}
