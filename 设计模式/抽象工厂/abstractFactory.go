package abstractFactory

import "fmt"

type Lunch interface {
	Cook()
}

type Rise struct {
}

func (r *Rise) Cook() {
	fmt.Println("it is rise")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("it is tomato")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type simpleLunchFactory struct {
}

func NewSimpleLunchFactory() LunchFactory {
	return &simpleLunchFactory{}
}

func (s *simpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s *simpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
