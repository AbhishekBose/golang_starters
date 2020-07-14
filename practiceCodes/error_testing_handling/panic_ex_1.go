package main

import "fmt"

func main() {
	fmt.Println("Starting")
	panic("Severe error occured")
	fmt.Println("ending program")
}
