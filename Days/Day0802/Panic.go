package main
import "os"
/*
panic 意味着有些出乎意料的错误发生。通常我们用它来表示程序正常运行中不应该出现的，
或者我们没有处理好的错误。
 */
func main() {
	panic("a problem")
	//panic 的一个基本用法就是在一个函数返回了错误值但是我们并不知道（或者不想）处理时终止运行。
	 _, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
/*
注意，不像在有些语言中使用异常处理错误，在 Go 中则习惯通过返回值来标示错误。
 */