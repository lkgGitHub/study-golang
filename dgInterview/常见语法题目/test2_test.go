package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// todo ???
// golang 的 for ... range 语法中，stu 变量会被复用，每次循环会将集合中的值复制给这个变量，
// 因此，会导致最后m中的map中储存的都是stus最后一个student的值。
func TestDemo12(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, ": ", *v)
	}
}

// todo 源码
// 第一个 只输出 i:10， 第二个i，会输出0~9
// 这个输出结果决定来自于调度器优先调度哪个G。
// 从runtime的源码可以看到，当创建一个G时，会优先放入到下一个调度的runnext字段上作为下一次优先调度的G。
// 因此，最先输出的是最后创建的G，也就是9.
func TestDemo13(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("==>i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func (p *Person) ShowA() {
	fmt.Println("person showA")
	p.ShowB() // person showB
}
func (p *Person) ShowB() {
	fmt.Println("person showB")
}

type Teacher struct {
	Person
}
type Person struct{}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

// 输出结果为person showA、person showB。golang 语言中没有继承概念，只有组合，也没有虚方法，更没有重载。
// 因此，*Teacher 的 ShowB 不会覆写被组合的 People 的方法。
func TestDemo14(t *testing.T) {
	teacher := Teacher{}
	teacher.ShowA() // person showA
	teacher.ShowB() // teacher showB
}

// 结果是随机执行。golang 在多个case 可读的时候会公平的选中一个执行。
func TestDemo15(t *testing.T) {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/*
输出：
10 1 2 3   3
20 0 3 2   2
2  0 2 2
1  1 3 4
*/
// defer 在定义的时候会计算好调用函数的参数，所以会优先输出10、20 两个参数。然后根据定义的顺序倒序执行。
func TestDemo16(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// make 在初始化切片时指定了长度，所以追加数据时会从len(s) 位置开始填充数据。
func TestDemo17(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //输出为  [0 0 0 0 0 1 2 3]
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

// 在执行 Get方法时可能被panic。
//
// 虽然有使用sync.Mutex做写锁，但是map是并发读写不安全的。map属于引用类型，并发读写时多个协程见是通过指针访问同一个地址，
// 即访问共享变量，此时同时读写资源存在竞争关系。会报错误信息:“fatal error: concurrent map read and map write”。
//
// 因此，在 Get 中也需要加锁，因为这里只是读，建议使用读写锁 sync.RWMutex。
func TestDemo18(t *testing.T) {
	u := UserAges{}
	u.ages = make(map[string]int)
	u.Add("a", 1)
	u.Add("b", 2)
	u.Add("c", 3)
	u.Get("a")
	u.Get("b")
	u.Get("c")

}

func TestDemo19(t *testing.T) {

}

type People2 interface {
	Speak(string) string
}

type Student2 struct{}

func (stu *Student2) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

//在 golang 语言中，Student 和 *Student 是两种类型，第一个是表示 Student 本身，第二个是指向 Student 的指针。
func TestDemo20(t *testing.T) {
	//var peo People2 = Student2{}  // 编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People类型
	var peo People2 = &Student2{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

type People3 interface {
	Show()
}

type Student3 struct{}

func (stu *Student3) Show() {

}
func live() People3 {
	var stu *Student3
	fmt.Printf("%T", stu) // *main.Student3
	return stu
}

// 跟上一题一样，不同的是*Student 的定义后本身没有初始化值，所以 *Student 是 nil的，
// 但是*Student 实现了 People 接口，接口不为 nil
func TestDemo21(t *testing.T) {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB") // 输出： BBBBBBB
	}
}
