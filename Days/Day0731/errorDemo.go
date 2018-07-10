package main
/*
Go 语言使用一个独立的·明确的返回值来传递错误信息的。这与使用异常的 Java 和 Ruby 以及
在 C 语言中经常见到的超重的单返回值/错误值相比，Go 语言的处理方式能清楚的知道哪个函数返回了错误，
并能像调用那些没有出错的函数一样调用。
 */
import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1,errors.New("can't work with 42")
	}
	return arg + 3, nil//返回错误值为 nil 代表没有错误。
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s",e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg,"can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7,42}{
		if r, e := f1(i); e != nil{
			fmt.Println("f1 failed:",e)
		}else {
			fmt.Println("f1 worked:",r)
		}
	}

	for _,i := range []int{7, 42}{
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:",e)
		}else {
			fmt.Println("f2 worked:",r)
		}
	}

	_, e := f2(42)
	if ae ,ok := e.(*argError); ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
