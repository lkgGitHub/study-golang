/*
   Template method模板方法：
        定义一个操作中的算法的骨架，而将一些具体步骤延迟到子类中。
		模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
 个人想法：与建造者：一个是行为型模式，一个是创建型模式
*/
package template

import "fmt"

type getFood interface {
	first()
	second()
	three()
}

type template struct {
	g getFood
}

func (b *template) getSomeFood() {
	b.g.first()
	b.g.second()
	b.g.three()
}

type bingA struct {
	template
}

func NewBingA() *bingA {
	return &bingA{}
}
func (b *bingA) first() {
	fmt.Println("打开冰箱")
}
func (b *bingA) second() {
	fmt.Println("拿出东西")
}
func (b *bingA) three() {
	fmt.Println("关闭冰箱")
}

type Guo struct {
	template
}

func NewGuo() *Guo {
	return &Guo{}
}
func (g *Guo) first() {
	fmt.Println("打开锅")
}

func (g *Guo) second() {
	fmt.Println("拿出东西锅")
}

func (g *Guo) three() {
	fmt.Println("关闭锅")
}
