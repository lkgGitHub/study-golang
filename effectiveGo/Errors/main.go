package main

import "os"

func main() {
	_, err := os.Open("/aaa")
	if err != nil {
		panic(err)
	}

}
