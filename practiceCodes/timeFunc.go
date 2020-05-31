package main
import "fmt"
import "time"

func Calculation(){
	for i := 0; i < 10000; i++ {
		//pass
	}
}

func main() {
	start := time.Now()
	Calculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Time taken to execute loop: %s\n",delta)
}
