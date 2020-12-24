package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type Param map[string]interface{}

type Show struct {
	Param
}

// new 关键字无法初始化 Show 结构体中的 Param 属性，所以直接对 s.Param 操作会出错。
// panic: assignment to entry in nil map
func TestDemo1(t *testing.T) {
	s := new(Show)
	s.Param["RMB"] = 10000
}

type student struct {
	Name string
}

func TestDemo2(t *testing.T) {
	//switch msg := v.(type) {
	//case *student, student:
	//	msg.Name
	//}
}


// 警告：结构字段 'name' 具有 'json' 标记，但未被导出
type PeopleLower struct {
	name string `json:"name"`
}
// 按照 golang 的语法，小写开头的方法、属性或 struct 是私有的，同样，在json 解码或转码的时候也无法上线私有属性的转换。
// 题目中是无法正常得到People的name值的。而且，私有属性name也不应该加json的标签。
func TestDemo3(t *testing.T) {
	js := `{
		"name":"11"
	}`
	var p PeopleLower
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	// 输出结果：people:  {}
	fmt.Println("PeopleLower: ", p)
}

type People struct {
	Name string
}
// 警告：占位符参数导致对 'String' 方法 (%v) 的递归调用
func (p *People) String() string {
	return "people"
	//return p.String()
	//return fmt.Sprintf("print: %v", p)
}

// todo ???
func TestDemo4(t *testing.T) {
	p := &People{}
	p.String()
}

// 在 golang 中 goroutine 的调度时间是不确定的，在题目中，第一个写 channel 的 goroutine 可能还未调用，
// 或已调用但没有写完时直接 close 管道，可能导致写失败，既然出现 panic 错误。
func TestDemo5(t *testing.T) {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

func TestDemo6(t *testing.T) {
	SetValue(3)
}
var value int32

func SetValue(delta int32) {
	//for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v+delta)) {
			println("value:", value)
			//break
		}
	//}
}