package main

import "fmt"
import "math"

const s string = "constant"
func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f :="short"
	//:= 语句是申明并初始化变量的简写
	fmt.Println(f)

	constants()
}

func constants(){
	fmt.Println(s)

	const n = 50000000
	//const 用于声明一个常量。
	const d = 2e3
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	//当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。举个例子，这里的 math.Sin函数需要一个 float64 的参数
}