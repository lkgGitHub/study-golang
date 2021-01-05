package facade

import "fmt"

/*
  Facade 外观模式：
        为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，
		这个接口使得这一子系统更加容易使用（投资：基金，股票，房产）
 个人想法：中介者模式、外观模式：每个对象都保存一份中介者对象，
        在和其他对象交互时，通过中介者来完成，外观模式：外观中保存了一堆对象，
		这些对象或者是组成某个子系统的，将其封装在外观对象中，给客户端一种只有一个对象的感觉，
		一个是结构型模式，一个是行为性模式
*/
type facade struct {
	one   funcOne
	two   funcTwo
	three funcThree
}

func NewFacade(i int, f float32, str string) *facade {
	return &facade{funcOne{i}, funcTwo{f}, funcThree{str}}
}

func (f facade) OutOne() {
	f.one.Out()
}

func (f facade) OutTwo() {
	f.two.Out()
	f.three.Out()
}

type funcOne struct {
	str int
}

func (f funcOne) Out() {
	fmt.Println("funcOne", f.str)
}

type funcTwo struct {
	str float32
}

func (f funcTwo) Out() {
	fmt.Println("funcTwo", f.str)
}

type funcThree struct {
	str string
}

func (f funcThree) Out() {
	fmt.Println("funcThree", f.str)
}
