package main

/*
  Decorator 装饰模式：
        动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更为灵活。
 个人想法：注意装饰器内部维护一个对象，所有装饰器的子类在操作时，必须调用父类的函数，一直从下到上再到下的感觉
*/

type AbsstractPerson interface {
	show()
}

type person struct {
	Name string
}

func (p *person) show() {
	println("姓名：", p.Name)
}

type Decorator struct {
	AbsstractPerson
}

func (d *Decorator) SetDecorator(component AbsstractPerson) {
	d.AbsstractPerson = component
}

func (d *Decorator) show() {
	if d.AbsstractPerson != nil {
		d.AbsstractPerson.show()
	}
}

type TShirts struct {
	Decorator
}

func (t *TShirts) show() {
	println("T恤")
}

type BigTrouser struct {
	Decorator
}

func (b *BigTrouser) show() {
	b.Decorator.show()
	println("大裤衩")
}

type Sneakers struct {
	Decorator
}

func (s *Sneakers) show() {
	s.Decorator.show()
	println("破冰鞋")
}
