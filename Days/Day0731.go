package main

import "fmt"

func main() {
	////闭包
	//nextInt := intSeq()
	//
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//
	//nextInts := intSeq()
	//fmt.Println(nextInts())
	////递归
	//fmt.Println(fact(5))
	////指针
	//i := 1
	//fmt.Println("initial:",i)
	//zeroval(i)
	//fmt.Println("zeroval:",i)
	//zeroptr(&i)
	//fmt.Println("zeroptr:",i)//通过 &i 语法来取得 i 的内存地址，例如一个变量i 的指针。
	//fmt.Println("pointer:",&i)//指针也是可以被打印的
	//// Go 的结构体 是各个字段字段的类型的集合。
	////使用这个语法创建了一个新的结构体元素。
	//fmt.Println(person{"Bob", 20})
	////你可以在初始化一个结构体元素时指定字段名字。
	//fmt.Println(person{name: "Alice", age: 30})
	////省略的字段将被初始化为零值。
	//fmt.Println(person{name: "Fred"})
	////& 前缀生成一个结构体指针。
	//fmt.Println(&person{name: "Ann", age: 40})
	////使用点来访问结构体字段。
	//s := person{name: "Sean", age: 50}
	//fmt.Println(s.name)
	////也可以对结构体指针使用. - 指针会被自动解引用。
	//sp := &s
	//fmt.Println(sp.age)
	////结构体是可变的。
	//sp.age = 51
	//fmt.Println(sp.age)
	// Go 支持在结构体类型中定义方法
	r := rect{width: 10, height: 5}

	fmt.Println("area:",r.area())
	fmt.Println("perim:",r.perim())

	rp := &r
	fmt.Println("area:",rp.area())
	fmt.Println("perim:",rp.perim())
}

type person struct {
	name string
	age int
}
// Go 支持在结构体类型中定义方法
type rect struct {
	width, height int
}
func (r *rect) area() int {
	return r.width * r.height
}
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Go 支持 指针，允许在程序中通过引用传递值或者数据结构。
func zeroval(ival int) {
	//zeroval 有一个 int 型参数，所以使用值传递。zeroval 将从调用它的那个函数中得到一个 ival形参的拷贝。
	ival = 0
}
func zeroptr (iptr *int){
	//有一和上面不同的 *int 参数，意味着它用了一个 int指针。函数体内的 *iptr 接着解引用 这个指针
	// ，从它内存地址得到这个地址对应的当前值。对一个解引用的指针赋值将会改变这个指针引用的真实地址的值。
	fmt.Println("iptr:",iptr)
	fmt.Println("*iptr:",*iptr)
	*iptr = 0
}
//递归
func fact(n int) int {
	if n == 0{
		return 0
	}
	return n+fact(n-1)
}

// Go 支持通过 闭包来使用 匿名函数。
func intSeq() func() int {
	i :=0
	return func() int {
		i += 1
		return i
	}
}
