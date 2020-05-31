package main

import "fmt"

func bubbleSort(sl []int) []int {
	n := len(sl)
	for n > 0 {
		fmt.Printf("N is %d\n", n)
		for i := 0; i < n-1; i++ {
			fmt.Println(sl[i])
			if sl[i] > sl[i+1] {
				fmt.Print("Need to swap\n")
				tmp := sl[i]
				sl[i] = sl[i+1]
				sl[i+1] = tmp
			}
		}
		n = n - 1
	}
	return sl
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(bubbleSort(arr))
}
