package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle   // 匿名成员后面不要跟,逗号
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 20
	w.Spokes = 20
	
	
	fmt.Println(w)
	
	
	// 通过字面值的方式
	w = Wheel{Circle{Point{1, 2}, 10}, 20}
	
	fmt.Println(w)
	
	// 或者这样
	w = Wheel {
		Circle: Circle{
			Point: Point{
				X: 10,
				Y: 10,
			},
			Radius: 40,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
	w.X = 42
	
	fmt.Printf("%#v\n", w)
}