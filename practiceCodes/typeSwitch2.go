package main

import "fmt"

type Simpler interface {
	Get() int
	Set(n int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	fmt.Println("Getting from Psimple")
	return p.i
}

func (p *Simple) Set(u int) {
	fmt.Println("Setting from Psimple")
	p.i = u

}

type RSimple struct {
	i int
}

func (p *RSimple) Get() int {
	fmt.Println("Getting from Rsimple")
	return p.i
}

func (p *RSimple) Set(u int) {
	fmt.Println("Setting into Rsimple")
	p.i = u * 10
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(50)
		return it.Get()
	default:
		return 99
	}
	return 0
}

func main() {
	var s Simple
	fmt.Println(fI(&s)) // &s is required because Get() is defined with the type pointer as receiver
	var r RSimple
	fmt.Println(fI(&r))
}
