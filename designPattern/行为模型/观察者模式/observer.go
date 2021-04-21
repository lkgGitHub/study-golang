/*
  Observer 观察者模式：
        定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。
		这个主题对象在状态发生改变时，会通知所有观察者对象，使它们能够自动更新自己。
 个人想法：抽象通知者依赖抽象观察者，具体观察者依赖具体通知者
*/
package observer

import "fmt"

type Subject interface {
	Notify()
	State() int
	SetState(int)
	AddCallFunc(*update)
	RemoveCallFunc(*update)
}

type update func(int)

type SubjectA struct {
	state int
	call  []*update
}

func (s *SubjectA) Notify() {
	for _, c := range s.call {
		(*c)(s.state)
	}
}

func (s *SubjectA) State() int {
	if s == nil {
		return 0
	}
	return s.state
}

func (s *SubjectA) SetState(i int) {
	if s == nil {
		return
	}
	s.state = i
}

func (s *SubjectA) AddCallFunc(u *update) {
	if s == nil {
		return
	}
	for _, c := range s.call {
		if c == u {
			return
		}
	}
	s.call = append(s.call, u)
}

func (s *SubjectA) RemoveCallFunc(u *update) {
	if s == nil {
		return
	}
	for i, c := range s.call {
		if c == u {
			s.call = append(s.call[:i], s.call[i+1:]...)
		}
	}
}

func NewSubjectA(s int) *SubjectA {
	return &SubjectA{
		state: s,
		call:  []*update{},
	}
}

type Observer interface {
	Update(int)
}

type ObserverA struct {
	s     Subject
	state int
}

func (o *ObserverA) Update(s int) {
	fmt.Println("ObserverA")
	fmt.Println(s)
	fmt.Println(o)
}

func NewObserverA(sa Subject, s int) *ObserverA {
	return &ObserverA{
		sa, s,
	}
}

type ObserverB struct {
	s     Subject
	state int
}

func (o *ObserverB) Update(s int) {
	if o == nil {
		return
	}
	fmt.Println("ObserverB")
	fmt.Println(s)
	fmt.Println(o)
}
func NewObserverB(sa Subject, s int) *ObserverB {
	return &ObserverB{sa, s}
}
