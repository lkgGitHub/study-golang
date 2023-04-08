package redigo_deme

import (
	"log"
	"testing"
	"time"
)

func TestSentinel(t *testing.T) {

	for {
		pool := newSentinelPool()
		conn := pool.Get()

		reply, err := conn.Do("get", "aaa")
		if err != nil {
			log.Println("error:", err.Error())
		}
		_ = conn.Close()

		if bs, ok := reply.([]byte); ok {
			log.Println("reply:", string(bs))
		}
		time.Sleep(time.Second * 3)

	}

}

func aaa() {

}
