package main

import "fmt"

func main() {
	fmt.Println("以下是数值的chan")
	ch := make(chan int, 3)
	ch <- 1
	close(ch)
	num, ok := <-ch
	fmt.Printf("读chan的协程结束，num=%v， ok=%v\n", num, ok)
	num1, ok1 := <-ch
	fmt.Printf("再读chan的协程结束，num=%v， ok=%v\n", num1, ok1)
	num2, ok2 := <-ch
	fmt.Printf("再再读chan的协程结束，num=%v， ok=%v\n", num2, ok2)

	fmt.Println("\n\n\n以下是字符串chan")
	cs := make(chan string, 3)
	cs <- "aaa"
	close(cs)
	str, ok := <-cs
	fmt.Printf("读chan的协程结束，str=%v， ok=%v\n", str, ok)
	str1, ok1 := <-cs
	fmt.Printf("再读chan的协程结束，str=%v， ok=%v\n", str1, ok1)
	str2, ok2 := <-cs
	fmt.Printf("再再读chan的协程结束，str=%v， ok=%v\n", str2, ok2)

	fmt.Println("\n\n\n以下是结构体chan")
	type MyStruct struct {
		Name string
	}
	cstruct := make(chan MyStruct, 3)
	cstruct <- MyStruct{Name: "haha"}
	close(cstruct)
	stru, ok := <-cstruct
	fmt.Printf("读chan的协程结束，stru=%v， ok=%v\n", stru, ok)
	stru1, ok1 := <-cs
	fmt.Printf("再读chan的协程结束，stru=%v， ok=%v\n", stru1, ok1)
	stru2, ok2 := <-cs
	fmt.Printf("再再读chan的协程结束，stru=%v， ok=%v\n", stru2, ok2)
}
