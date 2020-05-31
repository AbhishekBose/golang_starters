package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func main() {
	defer fmt.Println(Sum(1, 2))
	fmt.Println("Going to test sum function with defer")
}
