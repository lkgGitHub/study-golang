package interpreter

import "fmt"

/*
 Interpreter 解释器模式：
        给定一个语言，定义它的文法的一种表示，
		并定义一个解释器，这个解释器使用该表示来解释语言中的句子
*/

type Context struct {
	text string
}

// 抽象表达式
type IAbstractExpression interface {
	Interpret(ctx *Context)
}

// 终结符表达式
type TerminalExpression struct {
}

func (t *TerminalExpression) Interpret(ctx *Context) {
	ctx.text = ctx.text[:len(ctx.text)-1]
	fmt.Println(ctx)
}

// 非终结表达式
type NonTerminalExpression struct {
}

func (t *NonTerminalExpression) Interpret(ctx *Context) {
	ctx.text = ctx.text[:len(ctx.text)-1]
	fmt.Println(ctx)
}
