package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const key = "321423u9y8d2fwf9"

var storage map[string][][]byte
var once sync.Once

func sendToServer(identifier, content string) (int, error) {
	// 单例
	once.Do(func() {
		storage = make(map[string][][]byte)
	})

	encryptData, err := AesEncrypt([]byte(content), []byte(key))
	if err != nil {
		fmt.Println("AesEncrypt error:", err.Error())
		return 0, err
	}

	if _, ok := storage[identifier]; ok {
		storage[identifier] = append(storage[identifier], encryptData)
	} else {
		storage[identifier] = make([][]byte, 0)
		storage[identifier] = append(storage[identifier], encryptData)
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(serverCount), nil
}
