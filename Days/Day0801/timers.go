package main

import "time"
import "fmt"
func main() {
	//定时器表示在未来某一时刻的独立事件。你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。这里的定时器将等待 2 秒。
	timer1 := time.NewTimer(time.Second * 2)

	<- timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second *2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("Timer 2 stoped")
	}
}