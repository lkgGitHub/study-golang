package main
import "fmt"
import "sort"
func main() {
	//排序方法是正对内置数据类型的；这里是一个字符串的例子。注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值
	strs := []string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("String:",strs)

	ints := []int{2,8,6}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}