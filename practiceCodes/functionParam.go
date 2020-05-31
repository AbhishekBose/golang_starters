package main

import "fmt"

func main() {
	callback(1, Add)
}

func Add(x, y int) {
	fmt.Printf("Sum is::: %d", x+y)
}

func callback(x int, f func(x int, y int)) {
	f(x, 3)
}
