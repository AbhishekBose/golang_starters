package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {   //method calculating absolute value
	// fmt.Println(p.name)
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// func Abs(p Point) float64 {
// 	return math.Sqrt(p.x*p.x + p.y*p.y)
// }

type NamedPoint struct {
	Point
	name string
}

func main() {
	// n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	n := NamedPoint{Point{3, 4}, "Pythagoras"}
	// n:= Point{3,4}
	fmt.Println(n.Abs())
}
