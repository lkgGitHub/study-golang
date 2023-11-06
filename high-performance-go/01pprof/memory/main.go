package main

import (
	"math/rand"
	"strings"

	"github.com/pkg/profile"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 假设我们实现了这么一个程序，生成长度为 N 的随机字符串，拼接在一起。
func main() {
	// 我们使用一个易用性更强的库 pkg/profile 来采集性能数据，
	// pkg/profile 封装了 runtime/pprof 的接口，使用起来更简单。
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	concat(100)
}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

/*
两个字符串拼接时，相当于是产生新的字符串，如果当前的空间不足以容纳新的字符串，则会申请更大的空间，将新字符串完全拷贝过去，这消耗了 2 倍的内存空间。
使用 strings.Builder 替换 + 进行字符串拼接，将有效地降低内存消耗。
*/
func concatByBuilder(n int) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randomString(n))
	}
	return s.String()
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}
