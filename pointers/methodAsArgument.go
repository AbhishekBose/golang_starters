package main

import "fmt"

type T struct {
	a int
}

func (T) hello(message string) {
	fmt.Println("Hello!", message)

}

func callMethod(t T, method func(T, string)) {
	method(t, "a message")
}

func main() {
	t1 := T{10}
	callMethod(t1, T.hello)
}
