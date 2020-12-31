package adapter

/*
 Adapter 适配器模式：
        将一个类的接口转换成客户端希望的另一个接口。
		适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
 个人想法：代理和适配器：代理和代理的对象接口一致，客户端不知道代理对象，
        而适配器是客户端想要适配器的接口，适配器对象的接口和客户端想要的不一样，
		适配器将适配器对象的接口封装一下，改成客户端想要的接口
*/
import "fmt"

type Player interface {
	attack()
	defense()
}

type Forwards struct {
	name string
}

func (f *Forwards) attack() {
	println(f.name, "attacking")
}

func (f *Forwards) defense() {
	println(f.name, "defenseing")
}

func NewForwards(name string) Player {
	return &Forwards{name: name}
}

type Centers struct {
	name string
}

func (c *Centers) attack() {
	println(c.name, "attacking")
}

func (c *Centers) defense() {
	println(c.name, "defenseing")
}

func NewCenters(name string) Player {
	return &Centers{name: name}
}

type ForeignCenter struct {
	name string
}

func (f *ForeignCenter) attack(what string) {
	println(f.name, "attacking")
}

func (f *ForeignCenter) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

type Translator struct {
	f ForeignCenter
}

// 这是用户想要的接口
func (t *Translator) attack() {
	t.f.attack("Translator attack")
}

func (t *Translator) defense() {
	t.f.defense()
}

func NewTranslator(name string)Player{
	return &Translator{ForeignCenter{
		name: name,
	}}
}





