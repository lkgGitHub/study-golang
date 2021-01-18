package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	bg := BookGroup{}
	ba := Book{"a"}
	bg.add(ba)
	bb := Book{"b"}
	bg.add(bb)

	it := bg.createIterator()
	for b := it.first(); b != nil; b = it.next() {
		fmt.Println(b)
	}
}
