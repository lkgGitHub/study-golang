package main

import "fmt"
func main() {
	value := 1
	fmt.Println(value)     // prints 1
	{
		fmt.Println(value) // prints 1
		value := 2
		fmt.Println(value) // prints 2
	}
	fmt.Println(value)     // prints 1 (bad if you need 2)
}