package main
import "time"
import "fmt"
func main() {
	//打点器和定时器的机制有点相似：一个通道用来发送数据。这里我们在这个通道上使用内置的 range 来迭代值每隔500ms 发送一次的值。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C{
			fmt.Println("Tick at",t)
		}
	}()
	time.Sleep(time.Millisecond * 3000)
	ticker.Stop()
	fmt.Println("Ticeker stopped")
}