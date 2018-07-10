package main
import "fmt"
import "time"

func worker(done chan bool)  {
	fmt.Print("working...")
	fmt.Println(time.Second)
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true//发送一个值来通知我们已经完工啦。
}

func main() {
	done := make(chan bool,1)
	go worker(done)
	//程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done//如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。
}