package main
import "fmt"
//在这个例子中，我们将使用一个 jobs 通道来传递 main() 中 Go协程任务执行的结束信息到一个工作 Go 协程中。
// 当我们没有多余的任务给这个工作 Go 协程时，我们将 close 这个 jobs 通道。
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more{
				fmt.Println("received job", j)
			}else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job",j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}