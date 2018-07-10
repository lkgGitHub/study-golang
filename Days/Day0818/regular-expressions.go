package main

import "bytes"
import "fmt"
import "regexp"

// Go 提供内置的正则表达式。这里是 Go 中基本的正则相关功能的例子。
func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要 Compile 一个优化的 Regexp 结构体。
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("pench punch"))
	fmt.Println(r.FindAllStringSubmatch("peach punch pinch",1))
	fmt.Println(r.FindAllString("peach punch pinch",2))
	fmt.Print(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Print(r)

	fmt.Print(r.ReplaceAllString("a peach","<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in , bytes.ToUpper)
	fmt.Println(string(out))


}