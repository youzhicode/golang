package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X = factor
	p.Y = factor
}

func (p Path) Distance() float64 {
	var sum = 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{X: 4, Y: 6}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	p1 := Point{1, 2}
	p1.ScaleBy(5)

	p2 := Point{2, 3}
	pptr := &p2

	fmt.Printf("Point X=%v, Y=%v\n", p1.X, p1.Y)
	fmt.Printf("Point X=%v, Y=%v", pptr.X, pptr.Y)
}
