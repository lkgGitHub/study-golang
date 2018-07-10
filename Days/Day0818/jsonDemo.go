package main

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
	Page int
	Fruites []string
}
type Response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ :=json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(strB)

	slcD := []string{"apple","peach","pear"}
	slcB,_ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple":5,"lettuce":7}
	mapB,_ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	os.Exit(1)
}
