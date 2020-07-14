package main

import (
	"errors"
	"fmt"
)

func Sqrt(f float64) (float32, error) {
	if f < 0 {
		return 0, errors.New("Negative number encountered")
	} else {
		return 1, nil
	}
}

func main() {
	if f, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(f)
	}
}
