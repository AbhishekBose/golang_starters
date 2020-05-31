package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	a int
	b int
	int
	innerS
}

func main() {
	outer2 := outerS{6, 75, 60, innerS{5, 10}}
	fmt.Println(outer2)
	fmt.Println(outer2.in1)
}
