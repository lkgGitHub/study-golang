package observer

import "testing"

func TestObserver(t *testing.T) {
	s := NewSubjectA(1)
	oa := NewObserverA(s, 2)
	ob := NewObserverB(s, 2)
	var oafu update = oa.Update
	s.AddCallFunc(&oafu)
	var obfu update = ob.Update
	s.AddCallFunc(&obfu)
	println("=====================")

	s.Notify()
	println("=====================")

	s.SetState(3)
	s.Notify()

}
