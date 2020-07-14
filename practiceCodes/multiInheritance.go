package main
import "fmt"

type Camera struct{
	a string
}

func (c* Camera) TakeAPicture() string { //method of Camera
	return "Click"
}

type Phone struct {
	b string
 }


func (p* Phone) Call() string { //method of Camera
	return "Ring Ring"
}

type SmartPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(SmartPhone)
	fmt.Println("Our new SmartPhone exhibits multiple behaviors ...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
  }
