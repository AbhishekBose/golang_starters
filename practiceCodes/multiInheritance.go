package main

import "fmt"

type Camera struct{}

func (c *Camera) TakePicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring ring ring ring"
}

//multiple inheritance
type SmartPhone struct { //Can use method of both the structs
	Camera
	Phone
}

func main() {
	cp := new(SmartPhone)
	fmt.Println("It exhibits behaviour of a Camera: ", cp.TakePicture())
	fmt.Println(cp.Call())
}
