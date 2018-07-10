package main
//第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
import (
	"fmt"
	//"encoding/asn1"
)
//fmt 包实现了格式化 IO（输入/输出）的函数。
func main() {
//main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
	var a = 20  /* 声明实际变量 */
	var ip *int /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
	//fmt.Println("hello world")
	//fmt.Println(add(3,5))

}
func add(x int,y int) int{
	return x+y
}

