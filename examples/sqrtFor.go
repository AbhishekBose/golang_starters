// package main

// import "fmt"

// func Sqrt(x float64) float64 {
// 	// var z float64 = 1.0
// 	var prev float64 = 0.0
// 	for z := 1.0; z < 2.0; z++ {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Println(z)
// 		if z-prev == 0 {
// 			fmt.Println("Value reached")
// 			break
// 		}
// 		prev = z
// 	}
// 	return 0
// }

// func main() {
// 	returnValue := Sqrt(2)
// 	fmt.Println(returnValue)
// }
