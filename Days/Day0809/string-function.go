package main

import s "strings"
import "fmt"
/**
标准库的 strings 包提供了很多有用的字符串相关的函数。这里是一些用来让你对这个包有个初步了解的例子。
 */
var p = fmt.Println

func main() {
	p("Count: ",s.Count("test","t"))
	p("HasPrefix:",s.HasPrefix("test","te"))
	p("HasSuffix:",s.HasSuffix("test","st"))
	p("Index:",s.Index("test","e"))
	p("join:",s.Join([]string{"a","b"},"-"))
	p("repeat:",s.Repeat("a",5))
	p("Replace:",s.Replace("foo","o","0",-1))
	p("Replace:",s.Replace("foo","o","0",1))
	p("Split:",s.Split("a-b-c-d","-"))
	p("ToLower:",s.ToLower("TEST"))
	p("ToUpper:",s.ToUpper("test"))
	p("aaaaaaaaaaaaaaa")



}