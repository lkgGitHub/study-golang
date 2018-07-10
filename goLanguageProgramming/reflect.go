package main

import (
	"reflect"
	"fmt"
)

func main() {
	x := 3.14
	p := reflect.ValueOf(x)
	fmt.Println("type of p:",p.Type())
	fmt.Println("settability of P :",p.CanSet())

	p.SetFloat(3.5)
	fmt.Println(x)

	//v := p.Elem()
	//fmt.Println("settability of v: ", v.CanSet())
	//
	//v.SetFloat(3.5)
	//fmt.Println(v.Interface())
	//fmt.Println(x)
}
