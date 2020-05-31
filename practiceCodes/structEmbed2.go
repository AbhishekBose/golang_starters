package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += " " + s
}

func main() {
	c := new(Customer)
	c.Name = "Abhishek Bose"
	c.log = new(Log)
	c.log.msg = "Embedded struct"
	fmt.Println(*c.log)
	c.log.Add("Hello how are you")
	fmt.Println(*c.log)
}
