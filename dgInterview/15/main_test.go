package main

import (
	"fmt"
	"testing"
)

// 1. 写已经关闭的 chan
// panic: send on closed channel
func TestDemo(t *testing.T) {
	c := make(chan int, 3)
	close(c)
	c <- 1
}

// 2. 读已经关闭的 chan
func TestDemo2(t *testing.T) {
	fmt.Println("\n以下是数值的chan")
	ch := make(chan int, 3)
	ch <- 1
	close(ch)
	num, ok := <-ch
	fmt.Printf("读chan的协程结束，num=%v， ok=%v\n", num, ok)
	num1, ok1 := <-ch
	fmt.Printf("再读chan的协程结束，num=%v， ok=%v\n", num1, ok1)
	num2, ok2 := <-ch
	fmt.Printf("再再读chan的协程结束，num=%v， ok=%v\n", num2, ok2)

	fmt.Println("\n以下是字符串chan")
	cs := make(chan string, 3)
	cs <- "aaa"
	close(cs)
	str, ok := <-cs
	fmt.Printf("读chan的协程结束，str=%v， ok=%v\n", str, ok)
	str1, ok1 := <-cs
	println("str1 == \"\":", str1 == "")
	fmt.Printf("再读chan的协程结束，str=%v， ok=%v\n", str1, ok1)
	str2, ok2 := <-cs
	fmt.Printf("再再读chan的协程结束，str=%v， ok=%v\n", str2, ok2)

	fmt.Println("\n以下是结构体chan")
	type MyStruct struct {
		Name string
	}
	cstruct := make(chan MyStruct, 3)
	cstruct <- MyStruct{Name: "haha"}
	close(cstruct)
	stru, ok := <-cstruct
	fmt.Printf("读chan的协程结束，stru=%v， ok=%v\n", stru, ok)
	stru1, ok1 := <-cstruct
	println("stru1 == MyStruct{}:", stru1 == MyStruct{})
	fmt.Printf("再读chan的协程结束，stru=%v， ok=%v\n", stru1, ok1)
	stru2, ok2 := <-cstruct
	fmt.Printf("再再读chan的协程结束，stru=%v， ok=%v\n", stru2, ok2)
}

func TestDemo3(t *testing.T) {

}

func TestDemo4(t *testing.T) {

}
