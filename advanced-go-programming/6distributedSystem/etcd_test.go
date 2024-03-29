package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil { // handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	go func() { cli.Close() }()
	// 获取上下文，设置请求超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	fmt.Println(resp)

	_, err = cli.Get(ctx, "a")
	if err != nil {
		// with etcd clientv3 <= v3.3
		if err == context.Canceled {
			// grpc balancer calls 'Get' with an inflight client.Close
		} else if err == grpc.ErrClientConnClosing { // <= gRCP v1.7.x
			// grpc balancer calls 'Get' after client.Close.
		}
		// with etcd clientv3 >= v3.4
		if clientv3.IsConnCanceled(err) {
			// gRPC client connection is closed
		}
	}

	//m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	//if m == nil || err != nil {
	//	log.Printf("etcdsync.New failed")
	//	return
	//}
	//err = m.Lock()
	//if err != nil {
	//	log.Printf("etcdsync.Lock failed")
	//	return
	//}
	//
	//log.Printf("etcdsync.Lock OK")
	//log.Printf("Get the lock. Do something here.")
	//
	//err = m.Unlock()
	//if err != nil {
	//	log.Printf("etcdsync.Unlock failed")
	//} else {
	//	log.Printf("etcdsync.Unlock OK")
	//}
}
