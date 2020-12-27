package main

import (
	"fmt"
	"os"
	"time"
)

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	//s := new(Show)
	//s.Param = map[string]interface{}{"RMB": 1000}

	b := &Show{
		Param: map[string]interface{}{"RMB": 1000},
	}
	fmt.Printf("%v", b)
}

//func main() {
//	var x int
//	inc := func() int {
//		x++
//		return x
//	}
//	fmt.Println(func() (a, b int) {
//		return inc(), inc()
//	}())
//}

func test2() {
	abort := make(chan struct{})
	go func() {
		for {
			os.Stdin.Read(make([]byte, 1)) // read a single byte
			abort <- struct{}{}
		}
	}()

	bbb := make(chan string, 1)

	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case bbb <- "bbbb":
			b := <-bbb
			println(b)
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			break

		}

	}
	launch()
}

func launch() {
	println("launch")
}
