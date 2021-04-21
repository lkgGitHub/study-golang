package main

import "testing"

func TestDecorator(t *testing.T) {
	person := &person{"hcl"}
	person.show()
	println()
	ts := &TShirts{}
	ts.SetDecorator(person)
	ts.show()
	println()
	bt := &BigTrouser{}
	bt.SetDecorator(ts)
	bt.show()
	println()
	sk := &Sneakers{}
	sk.SetDecorator(bt)
	sk.show()
}
