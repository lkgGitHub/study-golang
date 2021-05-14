package main

import "fmt"

/*
golang程序变量会携带有一组校验数据，用来证明它的整个生命周期是否在运行时完全可知。如果变量通过了这些校验，它就可以在栈上分配。
否则就说它逃逸 了，必须在堆上分配。能引起变量逃逸到堆上的典型情况：

1. 在方法内把局部变量指针返回 局部变量原本应该在栈中分配，在栈中回收。但是由于返回时被外部引用，因此其生命周期大于栈，则溢出。
2. 发送指针或带有指针的值到 channel 中。 在编译时，是没有办法知道哪个 goroutine 会在 channel 上接收数据。
	所以编译器没法知道变量什么时候才会被释放。
3. 在一个切片上存储指针或带指针的值。 一个典型的例子就是 []*string 。
	这会导致切片的内容逃逸。尽管其后面的数组可能是在栈上分配的，但其引用的值一定是在堆上。
4. slice 的背后数组被重新分配了，因为 append 时可能会超出其容量( cap )。
	slice 初始化的地方在编译时是可以知道的，它最开始会在栈上分配。如果切片背后的存储要基于运行时的数据进行扩充，就会在堆上分配。
5. 在 interface 类型上调用方法。 在 interface 类型上调用方法都是动态调度的 —— 方法的真正实现只能在运行时知道。
	想像一个 io.Reader 类型的变量 r , 调用 r.Read(b) 会使得 r 的值和切片b 的背后存储都逃逸掉，所以会在堆上分配。

*/
type A struct {
	s string
}

// go build -gcflags=-m main.go
func main() {
	a := foo("hello")
	// a.s + " world" does not escape 说明 b 变量没有逃逸，因为它只在方法内存在，会在方法结束时被回收。
	b := a.s + " world"
	c := b + "!"
	// b + "!" escapes to heap 说明 c 变量逃逸，通过fmt.Println(a ...interface{})打印的变量，都会发生逃逸
	fmt.Println(c)
}

// 这是上面提到的 "在方法内把局部变量指针返回" 的情况
func foo(s string) *A {
	a := new(A)
	a.s = s
	// new(A) escapes to heap 说明 new(A) 逃逸了
	return a //返回局部变量a,在C语言中妥妥野指针，但在go则ok，但a会逃逸到堆
}
