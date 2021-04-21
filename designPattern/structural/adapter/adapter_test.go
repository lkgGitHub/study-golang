package adapter

import "testing"

func TestAdapter(t *testing.T) {
	pf := NewForwards("F")
	pc := NewCenters("C")
	pfc := NewTranslator("FC")

	pf.attack()
	pc.attack()
	pfc.attack()
	pfc.defense()

}