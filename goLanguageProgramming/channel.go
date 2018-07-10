package main

import (
	"fmt"
	"runtime"
)

func Count(ch chan int) {
	fmt.Println("Conting")
	ch <- 1
}

func main() {
	fmt.Println("cup:",runtime.NumCPU())
	runtime.Gosched()
	chs := make([]chan int,10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs{
		<- ch
	}

	ch2 := make(chan int,1)
	for false {
		select {
		case ch2 <- 0:
		case ch2 <- 1:
		}
		i := <- ch2
		fmt.Println("value received:",i)
}
//runtime.GOMAXPROCS(10)

}