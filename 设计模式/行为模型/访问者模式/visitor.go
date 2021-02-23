/*
 Visitor 访问者模式：
        表示一个作用于某对象结构中的各元素的操作，
		它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作
*/
package visitor

// 访问接口
type IVisitor interface {
	VisitConcreteElementA(ConcreteElementA)
	VisitConcreteElementB(ConcreteElementB)
}

// 具体访问者A
type ConcreteElementA struct {
	name string
}

func (c *ConcreteElementA) VisitConcreteElementA(ce ConcreteElementA) {
	println(ce.name, c.name)
	//ce.OperatorA()
}

type ConcreteElementB struct {
	name string
}
