package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestPutAndGet(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdEndpoint},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	defer func() { _ = cli.Close() }()

	// 获取上下文，设置请求超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := cli.Put(ctx, "sample_key", "sample_value")

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println(resp)
	_, err = cli.Get(context.Background(), "sample_key")
	if err != nil {
		t.Error(err.Error())
	}
}
