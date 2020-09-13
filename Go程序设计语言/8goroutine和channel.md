# CSP
goroutine和channel，其支持“顺序通信进程”(communicating sequential processes)或被简称为CSP。
# goroutines
goroutine泄露


# channels
1. 不带缓存的Channels
2. 串联的Channels（Pipeline）
3. 单方向的Channel
“类型chan<- int表示一个只发送int的channel，只能发送不能接收。相反，类型<-chan int表示一个只接收int的channel，只能接收不能发送。
（箭头<-和关键字chan的相对位置表明了channel的方向。）”
4. 带缓存的Channels

“channel的零值是nil”
“因为对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。”
//todo? 可以用nil来激活或者禁用case，来达成处理其它输入或输出事件时超时和取消的逻辑。

# 并发的循环
“这需要一种特殊的计数器，这个计数器需要在多个goroutine操作时做到安全并且提供在其减为零之前一直等待的一种方法。
这种计数类型被称为sync.WaitGroup”

# 多路复用

```go
select {
case <-ch1:
    // ...
case x := <-ch2:
    // ...use x...
case ch3 <- y:
    // ...
default:
    // ...
}
```
注意：
一个没有任何case的select语句写作select{}，会永远地等待下去。
如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会。

