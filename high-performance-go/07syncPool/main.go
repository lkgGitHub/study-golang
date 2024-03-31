package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return &Teacher{}
		},
	}
	teacher := &Teacher{
		Name: "ttt",
		Age:  18,
		Students: []*Student{
			{Name: "s1", Age: 6},
			{Name: "s2", Age: 7},
			{Name: "s3", Age: 8},
		},
	}
	pool.Put(teacher)

	t := pool.Get().(*Teacher)
	fmt.Printf("%+v \n", t) // &{Name:ttt Age:18 Students:[0xc000108488 0xc000108908 0xc000108d88]}

}

type Teacher struct {
	Name     string
	Age      uint8
	Students []*Student
}

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})
