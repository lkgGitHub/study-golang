package memento

import "testing"

func TestMemento(t *testing.T) {
	gr := GameRole{}
	gr.GetInitState()
	gr.StateDisplay()

	caretaker := RoleStateCaretaker{
		gr.SaveState(),
	}
	gr.Fight()
	gr.StateDisplay()
	gr.RecoveryState(caretaker.memento)
	gr.StateDisplay()

}
