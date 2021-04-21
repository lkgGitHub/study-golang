package bridge

/*
Bridge 桥接模式：
        将抽象部分与它的实现部分分离，使它们都可以独立地变化
 个人想法：组合/聚合复用原则

 */
type Phone struct {
	soft ISoftware
	name string
}

func (p *Phone) setSoft(soft ISoftware) {
	p.soft = soft
}

func (p *Phone) Run() {
	println(p.name)
	p.soft.Run()
}

type PhoneA struct {
	Phone
}

func NewPhoneA(name string) *PhoneA {
	return &PhoneA{Phone{name: name}}
}

type PhoneB struct {
	Phone
}

func NewPhoneB(name string) *PhoneB {
	return &PhoneB{
		Phone{name: name},
	}
}

type ISoftware interface {
	Run()
}

type Software struct {
	name string
}

type TSoftware struct{
	ISoftware
}

type SoftwareA struct {
	Software
}

func (s *SoftwareA) Run() {
	println(s.name)
}

type SoftwareB struct {
	Software
}

func (s *SoftwareB) Run() {
	if s == nil {
		return
	}
	println(s.name)
}