package chainofresponsibility

import "testing"

func TestChainofresponsibility(t *testing.T) {
	ca := NewConcreteHandlerA()
	cb := NewConcreteHandlerB()
	ca.SetSuccessor(cb)

	var req = [][]int{
		{4, constHandlerA},
		{11, constHandlerB},
		{0, constHandler},
	}
	for _, val := range req {
		if val[1] != ca.HandleRequest(val[0]) {
			t.Error("错误")
		}
	}

}
