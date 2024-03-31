package main

import (
	"bytes"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

// BenchmarkBufferWithPool-16      13758848                	88.28 ns/op
// BenchmarkBuffer-16        	   870776              		1416 ns/op
func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bu := bufferPool.Get().(*bytes.Buffer)
		bu.Write(data)
		bu.Reset()
		bufferPool.Put(bu)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var bu bytes.Buffer
		bu.Write(data)
	}
}
