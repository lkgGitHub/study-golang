package singletonMap

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	(*s)["this"] = "that"
	s2 := New()
	fmt.Println("This is ", (*s2)["this"])
}
