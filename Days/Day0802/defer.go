package main
import (
	"os"
	"fmt"
)
/*

Defer 被用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作。
 defer 用在像其他语言中的ensure 和 finally用到的地方
 */
//假设我们想要创建一个文件，向它进行写操作，然后在结束时关闭它。这里展示了如何通过 defer 来做到这一切。
func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
func createFile(p string) *os.File{
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f,"data")
}
func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}

