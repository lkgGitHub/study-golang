package main

import (
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	b, _:= ReadFile("/Users/mac/Documents/docs/16001864346450.md")
	println(string(b))

}


var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err!= nil{
		return nil, err
	}

	defer f.Close()
	return ioutil.ReadAll(f)
}
