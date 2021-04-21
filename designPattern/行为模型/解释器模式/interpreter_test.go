package interpreter

import "testing"

func TestInterpreter(t *testing.T) {
	context := Context{"abc"}
	var list []IAbstractExpression
	list = append(list, new(TerminalExpression))
	list = append(list, new(TerminalExpression))
	list = append(list, new(NonTerminalExpression))

	for _, val := range list {
		val.Interpret(&context)
	}
}
