package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	side float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)

	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)

	default:
		fmt.Printf("Unexpected type %T", t)
	}
	fmt.Println(areaIntf.Area())
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.side * ci.side * math.Pi
}
