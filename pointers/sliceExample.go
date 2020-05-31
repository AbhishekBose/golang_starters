package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	for i := 0; i < 6; i++ {
		arr1[i] = i
	}
	fmt.Println(("We are going to print the slice now"))
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
	slice2 := make([]int, 4, len(arr1)-2)
	slice2 = slice1[:]
	fmt.Println(slice2)
}
