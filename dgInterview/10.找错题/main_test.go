package main

import (
	"fmt"
	"testing"
)

// golang 中字符串是不能赋值 nil 的，也不能跟 nil 比较。
//func TestDemo(t *testing.T) {
//	var x string = nil
//	if x == nil {
//		x = "default"
//	}
//	fmt.Println(x)
//}

const (
	a = iota
	b = iota
)
const (
	c    = iota
	name = "name"
	//c    = iota
	d = iota
)

// 枚举类型
// const 中每新增一行常量声明将使 iota 计数一次( iota 可理解为 const 语句块中的行索引)。
func TestDemo2(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
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

// 上面的代码有严重的内存泄漏问题，出错的位置是 go fn(i)，
// 实际上代码执行后会启动 4 个协程，但是因为 ch 是非缓冲的，
// 只可能有一个协程写入成功。而其他三个协程会一直在后台等待写入。
func TestDemo3(t *testing.T) {
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

//golang 中的切片底层其实使用的是数组。当使用str1[1:] 使，str2 和 str1 底层共享一个数组，这回导致 str2[1] = "new" 语句影响 str1。
//而 append 会导致底层数组扩容，生成新的数组，因此追加数据后的 str2 不会影响 str1。
//但是为什么对 str2 复制后影响的确实 str1 的第三个元素呢？这是因为切片 str2 是从数组的第二个元素开始，str2 索引为 1 的元素对应的是 str1 索引为 2 的元素。
func TestDemo4(t *testing.T) {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1) // [a, b, new]
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1) // [a, b, new]
	fmt.Println(str2) // [b, new, z, x, y]
}

type Student struct {
	Name string
}

// 指针类型比较的是指针地址，非指针类型比较的是每个属性的值
func TestDemo5(t *testing.T) {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"}) // false
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})   // true
}

// 数组只能与相同纬度长度以及类型的其他数组比较。
// 切片之间不能直接比较。因此我们不能使用==操作符来判断两个slice是否含有全部相等元素。
// 不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等（[]byte），但是对于其他类型的slice，我们必须自己展开每个元素进行比较：
// slice唯一合法的比较操作是和nil比较
func TestDemo6(t *testing.T) {
	fmt.Println([...]string{"1"} == [...]string{"1"})
	//fmt.Println([]string{"1"} == []string{"1"})
}

type StudentAge struct {
	Age int
}

// golang中的map 通过key获取到的实际上是两个值，第一个是获取到的值，第二个是是否存在该key。
// 因此不能直接通过key来赋值对象。
func TestDemo7(t *testing.T) {
	kv := map[string]StudentAge{"menglu": {Age: 21}}
	//kv["menglu"].Age = 22 // error
	s := []StudentAge{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
}
