package template

import "testing"

func TestTemplate(t *testing.T) {
	b := NewBingA()
	b.g = b
	b.getSomeFood()
	println()
	g := NewGuo()
	g.g = g
	g.getSomeFood()
}
