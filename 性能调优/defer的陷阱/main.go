package main

import "fmt"

// 所以在defer语句中只能访问有名返回值，而不能直接访问匿名返回值。
func main() {
	fmt.Println("return:", deferCall())
	fmt.Println("return:", deferCall2())
}

// 匿名返回值是在return执行时被声明，
func deferCall() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer1:", i)  // 2
	}()
	defer func() {
		i++
		fmt.Println("defer2:", i)  // 1
	}()
	return i  // 0
}

// 有名返回值则是在函数声明的同时就已经被声明，
func deferCall2() (i int) {
	defer func() {
		i++
		fmt.Println("defer1:", i)  // 2
	}()
	defer func() {
		i++
		fmt.Println("defer2:", i) // 1
	}()
	return i // 2
}
