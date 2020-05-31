package main

import "fmt"

// make struct here
type anoStruct2 struct {
	x string
}
type anoStruct struct {
	val float32
	int
	anoStruct2
}

func main() {
	structEx := anoStruct(56.0, 12, "Hello")
	fmt.Println(structEx)
	// create struct using struct literal and print its field using fmt package
}
