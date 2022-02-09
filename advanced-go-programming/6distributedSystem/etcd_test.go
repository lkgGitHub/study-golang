package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func TestEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.2.242:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil { // handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	defer cli.Close()
	// 获取上下文，设置请求超时时间为5秒
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 设置key="/tizi365/url" 的值为 www.tizi365.com
	_, err = cli.Put(ctx, "/tizi365/url", "www.tizi365.com")
	if err != nil {
		log.Fatal(err)
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
