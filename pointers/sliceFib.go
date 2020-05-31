package main

import "fmt"

func fibarray(term int) []int {
	fibSlice := make([]int, term, term)
	for i := 0; i < term; i++ {
		if i <= 1 {
			fibSlice[i] = 1
		} else {
			fibSlice[i] = fibSlice[i-1] + fibSlice[i-2]
		}
	}
	return fibSlice
}

func main() {
	fmt.Println(fibarray(6))

}
