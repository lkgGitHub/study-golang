package main

import (
	"fmt"
)

const (
	c1 = iota
	c2
	c3
)

const c  = iota

func main() {
	i := 1
	j := 2
	i, j = j, i
	//fmt.Printf("i=%d,j=%d,c3=%d,c=%d", c1, c2, c3, c)
	//time.Sleep(10*time.Second)

	//str := "hello world"
	//for i,ch := range str{
	//	fmt.Println(i,ch)
	//}
	JLoop:
	for j := 0; j < 10; j++ {
		fmt.Print("j:",j,"\n")
		for i := 0; i < 10; i++ {
			fmt.Print("i:",i ,"\t")
			if i > 5 && j > 5{
				break JLoop
			}
		}
		fmt.Println()
		println()
	}
}
