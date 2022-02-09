package main

import (
	"io"
	"log"
	"net/http"

	"go.uber.org/ratelimit"
)

var rl ratelimit.Limiter

func sayHello(wr http.ResponseWriter, r *http.Request) {
	now := rl.Take()
	wr.WriteHeader(200)
	io.WriteString(wr, now.String())
}

func main() {
	rl = ratelimit.New(1100) // per second

	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
