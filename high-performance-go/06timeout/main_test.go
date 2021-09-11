package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

/*
最终程序中存在着 1002 个子协程，说明即使是函数执行完成，协程也没有正常退出。那越来越多的协程会残留在程序中，最终会导致内存耗尽（每个协程约占 2K 空间），程序崩溃。
done 是一个无缓冲区的 channel.当超时发生时，select 接收到 time.After 的超时信号就返回了，done 没有了接收方(receiver)，而 doBadThing 在执行 1s 后向 done 发送信号，由于没有接收者且无缓存区，发送者(sender)会一直阻塞，导致协程不能退出。
解决：即创建channel done 时，缓冲区设置为 1，即使没有接收方，发送方也不会发生阻塞。
*/
func TestBadTimeout(t *testing.T)  { test(t, doBadThing) }

// 协程数量下降为 2，创建的 1000 个子协程成功退出。
func TestBufferTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		timeoutWithBuffer(doBadThing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}
func TestGoodTimeout(t *testing.T) { test(t, doGoodThing) }

/*
还有一些更复杂的场景，例如将任务拆分为多段，只检测第一段是否超时，若没有超时，后续任务继续执行，超时则终止。


*/

func do2phases(phase1, done chan bool) {
	time.Sleep(time.Second) // 第 1 段
	select {
	case phase1 <- true:
	default:
		return
	}
	time.Sleep(time.Second) // 第 2 段
	done <- true
}

func timeoutFirstPhase() error {
	phase1 := make(chan bool)
	done := make(chan bool)
	go do2phases(phase1, done)
	select {
	case <-phase1:
		<-done
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func Test2phasesTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		timeoutFirstPhase()
	}
	time.Sleep(time.Second * 3)
	t.Log(runtime.NumGoroutine())
}
