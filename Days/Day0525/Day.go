package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Form:",r.Form)
	fmt.Println("path:",r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("r.Form[url_long]:",r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Print("key:",k)
		fmt.Println("===>val:", strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello world!")

}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.ListenAndServe(":9090",nil)
	//if err != nil {
	//	log.Fatal("ListenAndServer:",err)
	//}
}