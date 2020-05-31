package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	rect1 := Rectangle{5, 6}
	fmt.Println("Area of the rectangle is::: ", rect1.Area())
	fmt.Println("Perimeter of the rectangle is:::", rect1.Perimeter())
}
