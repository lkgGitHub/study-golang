package main
import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//通过键查询值
	val, err := client.Get("aaa").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("键golang的值为: ",val)

}

