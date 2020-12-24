package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

type Param map[string]interface{}

type Show struct {
	Param
}

// map 初始化，必须使用 make 后才能使用
// new 关键字无法初始化 Show 结构体中的 Param 属性，所以直接对 s.Param 操作会出错。
// panic: assignment to entry in nil map
func TestDemo1(t *testing.T) {
	s := new(Show) // 相当于var param map[string]interface{}， 只申明没有初始化
	s.Param["RMB"] = 10000
}

type student struct {
	Name string
	Age  int
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

// atomic.CompareAndSwapInt32 函数不需要循环调用。
func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			println("value:", value)
			break
		}
	}
}

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}
func (p *Project) exec(msgchan chan interface{}) {
	//defer p.deferError()  正确写法，应该放到这里。
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}
func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError() // 可能发生资源泄漏，在 'for' 循环中调用了 'defer'
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}
func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	// time.Sleep 的参数数值太大，超过了 1<<63 - 1 的限制
	time.Sleep(time.Second * 100000)
}

func TestDemo7(t *testing.T) {
	p := new(Project)
	p.Main()
}

func TestDemo8(t *testing.T) {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for {
			a := <-abc
			fmt.Println("a: ", a)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}

type Student struct {
	name string
}

// map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。
// 故如果需要修改map值，可以将map中的非指针类型value，修改为指针类型，比如使用map[string]*Student.
func TestDemo9(t *testing.T) {
	//m := map[string]Student{"people": {"zhoujielun"}}
	//m["people"].name = "wuyanzu"
}

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

// 依据4个goroutine的启动后执行效率，很可能打印111func4，
// 但其他的111func*也可能先执行，exec只会返回一条信息。
func TestDemo10(t *testing.T) {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}

// byte 其实被 alias 到 uint8 上了。
func BenchmarkDemo11(t *testing.B) {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ { // 条件 'i <= 255' 始终为 'true' . uint8 Range: 0 through 255.
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}
