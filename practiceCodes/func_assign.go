package main

import "fmt"

func main() {
	funcVar := func(x int, y int) int {
		return (x + y) * 2
	}
	fmt.Println(funcVar(3, 4))
}
