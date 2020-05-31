package main

import "fmt"

func Adder(a int) func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func main() {
	x := Adder(2)
	fmt.Println(x(3))
}
