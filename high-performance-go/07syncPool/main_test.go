package main

import (
	"encoding/json"
	"sync"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		_ = json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		studentPool.Put(stu)
	}
}

func unmarsh() {
	stu := &Student{}
	json.Unmarshal(buf, stu)
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}
