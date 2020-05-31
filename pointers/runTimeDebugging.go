package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	a := 25
	where()
	b := 100
	c := a + b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	where()
}
