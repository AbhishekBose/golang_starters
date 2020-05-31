package main
import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.WriteString("abc")
	b.WriteString("def")
	fmt.Fprintf(&b, " A number: %d, a string: %v\n", 10, "bird")
	b.WriteString("[DONE]")
	// Convert to a string and print it.
	fmt.Println(b.String())
}