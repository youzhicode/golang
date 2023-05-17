package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID			int
	Name		string
	Address 	string
	DoB			time.Time
	Position	string
	Salary 		int
	ManagerID	int
}

type Point struct {x, y int}

type address struct {
	hostname 	string
	port		int
}

/*
	定义了一个表示圆的结构体
*/
type Circle struct {
	X, Y, Radius int
}

/*
	定义了一个轮的结构体，不仅拥有Circle的成员，还增加了一个Spokes
*/
type Wheel struct {
	X, Y, Radius, Spokes int
}
// 随着项目的复杂度的增加，我们注意到上面两个结构体有许多相同和重复之处。为了便于维护，我们可以考虑把相同的属性独立
// 出来
/*

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point,
	Radius int
}

type Wheel Struct {
	MCircle Circle
	Spokes int
}

// 通过这样改动，条例变得清晰，但是访问变量变得悠长麻烦
var w Wheel
w.MCircle.Center.X = 8
w.MCircle.Center.Y = 8
w.Mcircle.Radius = 10
w.Mcircle.Spokes = 20

// GO语言中有一个特性，让我们只声明一个成员的数据类型也不用指出成员的名字
// 这类成员叫做匿名成员
// 匿名成员的类型必须是命名的类型或者指向一个命名类型的指针
// 使用匿名成员重新定义上面的结构体
type Circle struct {
	Point
	Radius int
}

type Whell struct {
	Circle
	Spokes int
}

// 得益于匿名的特性，我们访问叶子属性而不需要给出全路径
var w1 Wheel
w1.X = 8
w1.Y = 8
w1.Radius = 10
w1.Spokes = 20

*/

// 结构体可以作为函数的参数和返回值
func Scale(p Point, factor int) Point {
	return Point{p.x * factor, p.y * factor}
}

// 为了提高效率，一般使用结构体指针方式传入和返回
func Bouns(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// 如果需要修改结构体内部的成员，必须使用指针传入，因为在Go语言中，参数都是使用值传递的方式
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func main() {
	// 定义一个结构体的变量
	var dilbert Employee
	// 通过点操作符访问成员变量
	dilbert.Salary -= 5000
	// 通过指针访问成员变量
	position := &dilbert.Position
	*position = "Senior" + *position
	
	
	//点操作符结合指向结构体的指针
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "(proactive team player)"
	// 相当于(*employeeOfTheMonth).Position += "(proactive team player)"
	fmt.Println(employeeOfTheMonth)
	
	// 结构体也可以用结构体字面值来表示
	// 这种字面量，要求记住结构体的成员变量定义的顺序
//	p := Point{1, 2}
//	fmt.Println(p)
	// 更加常用的使用这种
	// anim := gif.GIF{LoopCount : 10}
	
	// 使用结构体作为参数，和返回值
	//fmt.Println(Scale(Point{1, 2}), 5)
	
	// 因为结构体通常通过指针来使用，可以用下面的写法来创建并初始化一个结构体变量
//	pp := &Point{1, 2}
	// 和下面的写法等价
//	pp2 := new(Point)
//	*pp2 = Point{1, 2}


	// 结构体的比较
	// 如果结构体的成员都是可以通过==比较的，那么结构体就可以通过==比较
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.x == q.x && p.y == q.y)
	fmt.Println(p == q)
	
	// 可比较的结构体类型跟其他可比较类型一样，可以用于map类型的key
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	
	// 结构体嵌入和匿名成员
	var w Whell
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	
	
	
}