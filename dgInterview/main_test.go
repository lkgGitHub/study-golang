package dgInterview

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var x int
	inc := func() int {
		x++
		return x
	}
	// result: 1 2
	fmt.Println(func() (a, b int) {
		return inc(), inc()
	}())
}
