package main

import "fmt"

type newType struct {
	name string
}

func (a *newType) String() string {
	return fmt.Sprintf(a.name)
}

func main() {
	x := newType{"hello"}
	y := newType{"world"}
	// x.name = "Abhishek"
	fmt.Println(x, y)
}
