package main
import "fmt"

func f(a [3]int) { fmt.Println(a) }   // accepts copy

func fp(a *[3]int) { fmt.Println(a) } // accepts pointer

func main() {
  var ar [3]int
  for i := 0; i < 3;i++ { ar[i] = i}
  f(ar) // passes a copy of ar
  fp(&ar) // passes a pointer to ar
}
