package main

import (
	"fmt"
	"time"
)

func main() {
	//forDemo()
	//ifElseDemo()
	//switchDemo()
	//arrayDemo()
	//sliceDemo()
	//mapDemo()
	//rangeDemo()
	//_,c := vals(2,3);fmt.Println(c)
	//fmt.Println(sum(1,2,3,4,5))
	//如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)重点是最后三个点
	nums := []int{1,2,3}
	fmt.Println(sum(nums...))
}
// Go 支持通过 闭包来使用 匿名函数。匿名函数在你想定义一个不需要命名的内联函数时是很实用的

//可变参数函数,这个函数使用任意数目的 int 作为参数
func sum(nums ...int) int{
	fmt.Print(nums)
	total := 0
	for _,num := range nums{
		total +=num
	}
	return total
}

//返回多个值
func vals(a int, b int)(string,int){
	return "a+b=",a+b
}

func rangeDemo() {
	nums := []int{1,2,3}
	sum := 0
	for _, num := range nums{//空值定义符_
		sum += num
	}
	fmt.Println("sum:",sum)

	for i, num := range nums{
		if num == 3 {
			fmt.Println("index:",i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i,c := range "go"{
		fmt.Println(i,c)
	}
}
// map 是 Go 内置关联数据类型（在一些其他的语言中称为哈希 或者字典 ）
func mapDemo() {
	//要创建一个空 map，需要使用内建的 make:make(map[key-type]val-type)
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	v1 :=m["k1"]
	fmt.Println("map:",v1)

	fmt.Println("len",len(m))

	delete(m,"k2")
	fmt.Println("map:",m)

	prs :=m["k2"]
	fmt.Println("prs",prs)

	n := map[string]int{"foo":1,"bar":2}
	fmt.Println(n)

}

// Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口
func sliceDemo() {
	s := make([]string, 3)
	fmt.Println("emp:",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	c := make([]string,len(s))
	copy(c, s)
	fmt.Println("copy:",c)

	//Slice 支持通过 slice[low:high] 语法进行“切片”操作,（不包含最高位）
	l := s[2:3]
	fmt.Println("sl1:",l)

	l = s[:3]
	fmt.Println("sl1:",l)

	l = s[2:]
	fmt.Println("sl1:",l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:",t)

	twoD := make([][]int,3)
	for i := 0; i < 3; i++{
		innerLen :=i + 1
		twoD[i] = make([]int,innerLen)
		for j :=0; j < innerLen; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:",twoD)

}
//数组
func arrayDemo()  {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println("set",a)
	fmt.Println("get",a[4])
	fmt.Println("len",len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	var twoD [2][3]int
	for i :=0; i < 2; i++{
		for j :=0; j < 3; j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println(twoD)
}
// switch ，方便的条件分支语句
func switchDemo() {
	i := 2
	fmt.Print("write",i,"as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	fmt.Println("============",time.Now(),"============",time.Now().Weekday())

}
//if 和 else 分支结构,Go 里没有三目运算符
func ifElseDemo(){
	if 7%2 == 0 {
		fmt.Println("7 is even")
	}else {
		fmt.Println("7 is odd")
	}

}
//for 是 Go 中唯一的循环结构。
func forDemo() {
	i := 1
	for i <= 3{
		fmt.Println(i)
		i++
	}
	for j := 7; j<=9; j++{
		fmt.Println(j)
	}
	for{
		fmt.Println("loop")
		break
	}
}
