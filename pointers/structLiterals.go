package main

import "fmt"

type Details struct {
	name string
	age int
	location string
}

var userDetails map[string]Details

func main() {
	userDetails = make(map[string]Details)
	userDetails["user1"] = Details{"ross",30,"Delhi"}
	fmt.Println(userDetails["user1"].name)
}