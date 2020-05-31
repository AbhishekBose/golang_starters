package main

import "fmt"

func main() {
	var maplit map[string]int
	maplit = map[string]int{"one": 1, "two": 2, "three": 3}
	mapcreated := make(map[string]float64)
	mapcreated["one_key"] = 12.4
	fmt.Println((maplit))
	fmt.Println(mapcreated)
	fmt.Println(mapcreated["one"])
	fmt.Println(maplit["one"])
	value, present := mapcreated["one_key"]
	fmt.Println(value)
	fmt.Println(present)
	delete(maplit, "three")
	fmt.Println(maplit)
}
