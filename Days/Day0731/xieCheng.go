package main
//Go 协程 在执行上来说是轻量级的线程。
import "fmt"

func f(from string)  {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":",i)
	}
}

func main() {
	//f("direct")

	go f("goroutine")

	go f("going")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)//这里的 Scanln 代码需要我们在程序退出前按下任意键结束。
	fmt.Println("done")
}
