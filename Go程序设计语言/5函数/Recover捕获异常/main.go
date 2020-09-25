package main

import (
	"fmt"
	"golang.org/x/net/html"
)

func main() {

}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct {}
	defer func() {
		switch p:=recover();p {
		case nil:
			println("no panic")
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if
	})
}
