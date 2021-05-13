/*
 Strategy 策略模式：
        它定义了算法家族，分别封装起来，让它们可以相互替换，
		此模式让算法的变化，不会影响到使用算法的客户。
 个人想法：UML图很相似，策略模式是用在对多个做同样事情（统一接口）的类对象的选择上，
         而状态模式是：将对某个事情的处理过程抽象成接口和实现类的形式，
		由context保存一份state，在state实现类处理事情时，修改状态传递给context，
		由context继续传递到下一个状态处理中，
*/
package strategy

import "errors"

type CashSuper interface {
	acceptCash(memory float32) float32
}

type CashNormal struct {

}
func (c *CashNormal) acceptCash(memory float32) float32 {
	return memory
}

type CashRebate struct {
	memoryRebate float32
}
func (c *CashRebate) acceptCash(memory float32) float32 {
	return c.memoryRebate * memory
}

type CashReturn struct {
	memoryCondition float32
	memoryReturn float32
}
func (c *CashReturn) acceptCash(memory float32) float32 {
	if memory > c.memoryCondition {
		return memory - float32(int(memory/c.memoryCondition*c.memoryCondition))
	}else{
		return memory
	}
}

type Context struct {
	CashSuper
}

func NewCashContext(str string) (cash *Context, err error)  {
	cash = &Context{}
	switch str {
	case "正常收费":
		cash.CashSuper = &CashNormal{}
	case "满300返100":
		cash.CashSuper = &CashReturn{300, 100}
	case "打8折":
		cash.CashSuper = &CashRebate{0.8}
	default:
		err = errors.New("no match")
	}
	return cash, err
}




