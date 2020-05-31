package main

import "fmt"

type flt func(int) bool
type slice_split func([]int) ([]int, []int)

func isOdd(x int) bool {
	if x%2 == 0 {
		return false
	} else {
		return true
	}
}
func filter(f flt) slice_split {
	return func(s []int) (yes, no []int) {
		for _, val := range s {
			if f(val) {
				yes = append(yes, val)
			} else {
				no = append(no, val)
			}
		}
		return
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("s = ", s)
	odd_even_function := filter(isOdd)
	odd, even := odd_even_function(s)
	fmt.Println("odd = ", odd)
	fmt.Println("even = ", even)
}
