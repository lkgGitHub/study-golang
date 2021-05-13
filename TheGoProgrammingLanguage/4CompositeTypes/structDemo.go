package main

import "fmt"

// 结构体嵌入和匿名成员

func main() {
	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)

	w = Wheel{
		Circle: Circle{
			Point: Point {X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)

	w.X = 55
	fmt.Printf("%#v\n", w)

}

type Point struct{ X, Y int }

// 一个Circle代表的圆形类型包含了标准圆心的X和Y坐标信息，和一个Radius表示的半径信息
type Circle struct {
	Point
	Radius int
}

// 一个Wheel轮形除了包含Circle类型所有的全部成员外，还增加了Spokes表示径向辐条的数量。
type Wheel struct {
	Circle
	Spokes int
}





