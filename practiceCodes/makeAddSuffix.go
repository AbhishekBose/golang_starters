package main

import (
	"fmt"
	"strings"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addJpeg := MakeAddSuffix(".jpg")
	addBmp := MakeAddSuffix(".bmp")
	fmt.Println(addJpeg("Hello"))
	fmt.Println(addBmp("Hello"))
}
