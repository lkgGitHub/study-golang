package main

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags '-s -w -extldflags "-static"' -o etcdTest main.go
import (
	"context"
	"flag"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var etcdEndpoint = "192.168.2.242:22379"

func main() {
	flag.StringVar(&etcdEndpoint, "etcd", etcdEndpoint, "etcd endpoint")
	flag.Parse()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdEndpoint},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go func() { _ = cli.Close() }()

	// 获取上下文，设置请求超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	tick := time.NewTicker(time.Second)

	for {
		putResponse, err := cli.Put(ctx, "sample_key", "sample_value")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("putResponse:", putResponse)

		getResponse, err := cli.Get(ctx, "sample_key")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("getResponse:", getResponse)
		<-tick.C
	}
}
