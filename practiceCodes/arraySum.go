package main

import "fmt"

func Sum(arr *[3]int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(Sum(&arr))

}
