package flyweight

import "testing"

func TestFlyWeight(t *testing.T) {
	ff := NewFlyWeightFactory()

	fya := ff.FlyWeight("a")
	fya.Operation(1)

	fyb := ff.FlyWeight("b")
	fyb.Operation(2)

	fyc := ff.FlyWeight("c")
	fyc.Operation(3)

	fyd := ff.FlyWeight("d")
	fyd.Operation(4)

	fyu := ff.FlyWeight("u")
	fyu.Operation(5)
}
