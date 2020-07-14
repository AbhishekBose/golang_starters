package main

import "fmt"

type Square struct {
	side float32
}

type Triangle struct {
	base   float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Perimeter() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return 4 * sq.side
}

func (tr *Triangle) Area() float32 {
	return 0.5 * tr.base * tr.height
}

func main() {
	var areaIntf AreaInterface
	var periIntf PeriInterface

	s1 := new(Square)
	s1.side = 10
	t1 := new(Triangle)
	t1.base = 10
	t1.height = 10
	periIntf = s1
	fmt.Println("Perimeter of square is:: ", periIntf.Perimeter())

	areaIntf = s1
	fmt.Println("Area of square is:: ", areaIntf.Area())

	areaIntf = t1
	// var perint PeriInterface

	fmt.Println("Area of traingle is:: ", areaIntf.Area())

}
