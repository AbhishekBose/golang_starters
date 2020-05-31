package main

import "fmt"

type Point struct {
	X, Y float64
}

func (a Point) Abs() float64 {
	return 9.0
}

func (p *Point) Scale(s float64) *Point {
	p.X = s * p.X
	p.Y = s * p.Y
	return p
}

func main() {
	p1 := Point{1.0, 2.0}
	// p1.Scale(5)
	fmt.Println(*p1.Scale(5))
}
